using System;
using System.Collections.Concurrent;
using System.Linq;
using System.Net;
using System.Net.Sockets;
using System.Threading;
using System.Threading.Tasks;
using Relax.MmoGame.Common;
using Relax.MmoGame.Common.Extensions;

namespace Relax.MmoGame.Server
{
    public class GameServer : IDisposable, IAsyncDisposable
    {
        private const byte WorldSize = 25;
        private const string Address = "127.0.0.1";

        private readonly TcpListener server;

        private CancellationTokenSource _serverCancellationTokenSource = new();

        #region Players

        private readonly ConcurrentDictionary<byte, PlayerProcessor> _playerProcessors = new();

        private readonly PlayerWatcher _playerWatcher;
        private readonly byte[] _writeBuffer = new byte[Consts.MaxServerWrite];

        private readonly System.Timers.Timer _timer = new(1000 / Consts.Fps);

        #endregion

        public GameServer(int port)
        {
            server = new TcpListener(IPAddress.Parse(Address), port);
            _playerWatcher = new PlayerWatcher(WorldSize);

            _timer.AutoReset = true;
            _timer.Elapsed += (_, _) => SendPositionToAll();
        }

        public async Task Start(CancellationToken cancellationToken)
        {
            _serverCancellationTokenSource = CancellationTokenSource.CreateLinkedTokenSource(cancellationToken);
            var innerCancellationToken = _serverCancellationTokenSource.Token;

            try
            {
                server.Start();

                EventLoop(innerCancellationToken).RunBackground(innerCancellationToken);

                while (!innerCancellationToken.IsCancellationRequested)
                {
                    var client = await server.AcceptTcpClientAsync();
                    _timer.Enabled = true;

                    var playerProcessor = new PlayerProcessor(_playerWatcher);
                    playerProcessor.Run(client, innerCancellationToken);
                    _playerProcessors.TryAdd(playerProcessor.Player.PlayerId, playerProcessor);

                    Console.WriteLine($"Client {playerProcessor.Player.PlayerId} accepted!");
                }
            }
            catch (Exception ex)
            {
                Console.WriteLine(ex.Message);
            }
        }

        private async Task EventLoop(CancellationToken cancellationToken)
        {
            while (!cancellationToken.IsCancellationRequested)
            {
                await Task.Delay(100, cancellationToken);

                foreach (var processor in _playerProcessors.Values)
                {
                    if (processor.CanRead)
                    {
                        processor.ReceivePackages().RunBackground(cancellationToken);
                    }

                    if (processor.CanWrite)
                    {
                        processor.SendPackages().RunBackground(cancellationToken);
                    }
                }
            }
        }

        private void SendPositionToAll()
        {
            if (!_playerProcessors.Any())
            {
                return;
            }

            var players = _playerProcessors.Values
                .Where(processor => processor.Connected)
                .Select(processor => processor.Player).ToArray();

            var bytes = new DrawPackage {Players = players}.Serialize(_writeBuffer);

            foreach (var processor in _playerProcessors.Values)
            {
                if (processor.Connected)
                {
                    processor.AddPackage(bytes);
                }
                else
                {
                    _playerProcessors.TryRemove(processor.Player.PlayerId, out _);
                }
            }
        }

        public void Dispose()
        {
            _timer?.Dispose();

            foreach (var processor in _playerProcessors.Values)
            {
                processor?.Dispose();
            }

            server?.Stop();
            _serverCancellationTokenSource?.Dispose();
        }

        public ValueTask DisposeAsync()
        {
            Dispose();

            return ValueTask.CompletedTask;
        }
    }
}