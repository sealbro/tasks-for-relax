using System;
using System.Collections.Concurrent;
using System.Net.Sockets;
using System.Threading;
using System.Threading.Tasks;
using Relax.MmoGame.Common;
using Relax.MmoGame.Common.Extensions;

namespace Relax.MmoGame.Client
{
    public class GameClient : IDisposable, IAsyncDisposable, IClientService
    {
        private readonly string _address;
        private readonly int _port;
        private readonly TcpClient client;

        private readonly byte[] _readBuffer = new byte[Consts.MaxClientRead];
        private readonly byte[] _writeBuffer = new byte[Consts.MaxClientWrite];
        private readonly ConcurrentQueue<byte[]> _packagesToSend = new();
        private CancellationTokenSource _cancellationTokenSource = new();

        #region Player

        private bool _registered = false;
        private byte _playerId = 0;

        private readonly System.Timers.Timer _timer = new(1000);

        #endregion

        private readonly ServerPackageProcessor _serverPackageProcessor;
        private WorldMap _map;

        public GameClient(string address, int port)
        {
            _address = address;
            _port = port;
            client = new TcpClient(_address, _port);

            _serverPackageProcessor = new ServerPackageProcessor(this);

            _timer.AutoReset = true;
            _timer.Elapsed += (_, _) => MovePlayer();
        }

        public async Task Start(CancellationToken cancellationToken)
        {
            _cancellationTokenSource = CancellationTokenSource.CreateLinkedTokenSource(cancellationToken);
            var innerCancellationToken = _cancellationTokenSource.Token;

            try
            {
                var stream = client.GetStream();

                await EventLoop(stream, innerCancellationToken);
            }
            catch (Exception ex)
            {
                Console.WriteLine(ex.Message);
            }
        }

        #region IClientService

        public void Registered(RegisteredPackage package)
        {
            _registered = true;
            _timer.Enabled = _registered;
            _playerId = package.PlayerId;

            _map = new WorldMap(package.WorldSize);
        }

        public void Exit()
        {
            _registered = false;
            _cancellationTokenSource.Cancel();
        }

        public void Draw(DrawPackage package)
        {
            if (!_registered)
            {
                _timer.Enabled = false;
                return;
            }

            _map.Draw(package.Players);
        }

        #endregion

        private void MovePlayer()
        {
            if (!_registered) return;

            var bytes = new MovePackage
            {
                Direction = BotPlayer.GetDirection(), PlayerId = _playerId
            }.Serialize(_writeBuffer);

            _packagesToSend.Enqueue(bytes);
        }

        private async Task EventLoop(NetworkStream stream, CancellationToken cancellationToken)
        {
            while (!cancellationToken.IsCancellationRequested)
            {
                await Task.Delay(100, cancellationToken);

                if (stream.CanWrite)
                {
                    WritePackages(stream, cancellationToken).RunBackground(cancellationToken);
                }

                if (stream.CanRead && stream.DataAvailable)
                {
                    ReadPackages(stream, cancellationToken).RunBackground(cancellationToken);
                }
            }
        }

        private async Task ReadPackages(NetworkStream stream, CancellationToken cancellationToken)
        {
            await stream.ReadAsync(_readBuffer, cancellationToken);

            var command = _serverPackageProcessor.ReceiveProcess(_readBuffer);

            Console.WriteLine(DateTime.Now.ToString("hh:mm:ss.fff") + $": {_playerId} => {command}");
        }

        private async Task WritePackages(NetworkStream stream, CancellationToken cancellationToken)
        {
            while (_packagesToSend.TryDequeue(out var bytes))
            {
                await stream.WriteAsync(bytes, cancellationToken);
            }
        }

        public void Dispose()
        {
            _timer?.Dispose();
            client?.Close();
            _cancellationTokenSource?.Dispose();
        }

        public ValueTask DisposeAsync()
        {
            Dispose();

            return ValueTask.CompletedTask;
        }
    }
}