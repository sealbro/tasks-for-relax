using System;
using System.Collections.Concurrent;
using System.Linq;
using System.Net;
using System.Net.Sockets;
using System.Threading;
using System.Threading.Tasks;
using Relax.MmoGame.Common;

namespace Relax.MmoGame.Server
{
    public class GameServer : IDisposable, IAsyncDisposable
    {
        private const int Fps = 2;
        private const byte WorldSize = 10;
        private const string Address = "127.0.0.1";

        private readonly TcpListener server;

        private CancellationTokenSource _serverCancellationTokenSource = new();

        #region Players

        private readonly ConcurrentDictionary<byte, PlayerProcessor> _playerProcessors = new();

        private readonly PlayerWatcher _playerWatcher;
        private readonly byte[] _writeBuffer = new byte[500];

        #endregion

        public GameServer(int port)
        {
            server = new TcpListener(IPAddress.Parse(Address), port);
            _playerWatcher = new PlayerWatcher(WorldSize);
        }

        public async Task Start(CancellationToken cancellationToken)
        {
            _serverCancellationTokenSource = CancellationTokenSource.CreateLinkedTokenSource(cancellationToken);
            var innerCancellationToken = _serverCancellationTokenSource.Token;

            try
            {
                server.Start();

                SendPositionToAll(innerCancellationToken);

                while (!innerCancellationToken.IsCancellationRequested)
                {
                    var client = await server.AcceptTcpClientAsync();

                    Console.WriteLine("Client accepted!");

                    var playerProcessor = new PlayerProcessor(_playerWatcher);
                    playerProcessor.Run(client, innerCancellationToken);
                    _playerProcessors.TryAdd(playerProcessor.Player.PlayerId, playerProcessor);
                }
            }
            catch (Exception ex)
            {
                Console.WriteLine(ex.Message);
            }
        }

        private async void SendPositionToAll(CancellationToken cancellationToken)
        {
            await Task.Run(async () =>
            {
                while (!cancellationToken.IsCancellationRequested)
                {
                    await Task.Delay(1000 / Fps, cancellationToken);

                    if (!_playerProcessors.Any())
                    {
                        continue;
                    }

                    var players = _playerProcessors.Values
                        .Where(processor => processor.Connected)
                        .Select(processor => processor.Player).ToArray();

                    var bytes = new DrawPackage {Players = players}.Serialize(_writeBuffer);

                    foreach (var processor in _playerProcessors.Values)
                    {
                        if (!processor.Connected)
                        {
                            _playerProcessors.TryRemove(processor.Player.PlayerId, out _);
                        }

                        processor.AddPackage(bytes);
                    }
                }
            }, cancellationToken);
        }

        public void Dispose()
        {
            foreach (var processor in _playerProcessors.Values)
            {
                processor.Dispose();
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