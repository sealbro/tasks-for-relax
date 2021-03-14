using System;
using System.Collections.Concurrent;
using System.Net.Sockets;
using System.Threading;
using System.Threading.Tasks;
using Relax.MmoGame.Common;

namespace Relax.MmoGame.Client
{
    public class GameClient : IDisposable, IAsyncDisposable, IClientService
    {
        private readonly string _address;
        private readonly int _port;
        private readonly TcpClient client;

        private readonly byte[] _readBuffer = new byte[50];
        private readonly byte[] _writeBuffer = new byte[50];
        private readonly ConcurrentQueue<byte[]> _packagesToSend = new();
        private CancellationTokenSource _cancellationTokenSource = new();

        #region Player

        private bool _registered = false;
        private byte _playerId = 0;

        #endregion

        private readonly ServerPackageProcessor _serverPackageProcessor;
        private WorldMap _map;

        public GameClient(string address, int port)
        {
            _address = address;
            _port = port;
            client = new TcpClient(_address, _port);

            _serverPackageProcessor = new ServerPackageProcessor(this);
        }

        public async Task Start(CancellationToken cancellationToken)
        {
            _cancellationTokenSource = CancellationTokenSource.CreateLinkedTokenSource(cancellationToken);
            var innerCancellationToken = _cancellationTokenSource.Token;

            try
            {
                var stream = client.GetStream();

                await Task.WhenAll(
                    ReceivePackage(stream, innerCancellationToken),
                    SendPackage(stream, innerCancellationToken));
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
            _playerId = package.PlayerId;

            _map = new WorldMap(package.WorldSize);

            SendDirection(_playerId, _cancellationTokenSource.Token);
        }

        public void Exit()
        {
            _registered = false;
            _cancellationTokenSource.Cancel();
        }

        public void Draw(DrawPackage package)
        {
            if (!_registered) return;

            _map.Draw(package.Players);
        }

        #endregion

        private async void SendDirection(byte playerId, CancellationToken cancellationToken)
        {
            await Task.Run(async () =>
            {
                while (!cancellationToken.IsCancellationRequested)
                {
                    await Task.Delay(1000, cancellationToken);

                    if (!_registered) continue;

                    var bytes = new MovePackage
                    {
                        Direction = BotPlayer.GetDirection(), PlayerId = playerId
                    }.Serialize(_writeBuffer);

                    _packagesToSend.Enqueue(bytes);
                }
            }, cancellationToken);
        }

        private async Task SendPackage(NetworkStream stream, CancellationToken cancellationToken)
        {
            // await stream.WriteAsync(new ConnectPackage().Serialize(_writeBuffer), cancellationToken);

            while (!cancellationToken.IsCancellationRequested && stream.CanWrite)
            {
                await Task.Delay(100, cancellationToken);

                while (_packagesToSend.TryDequeue(out var bytes))
                {
                    await stream.WriteAsync(bytes, cancellationToken);
                }
            }
        }

        private async Task ReceivePackage(NetworkStream stream, CancellationToken cancellationToken)
        {
            while (!cancellationToken.IsCancellationRequested && stream.CanRead)
            {
                await Task.Delay(100, cancellationToken);

                if (!stream.DataAvailable) continue;

                await stream.ReadAsync(_readBuffer, cancellationToken);

                var command = _serverPackageProcessor.ReceiveProcess(_readBuffer);

                Console.WriteLine(DateTime.Now.ToString("hh:mm:ss.fff") + $": {_playerId} => {command}");
            }
        }

        public void Dispose()
        {
            client?.Close();
            _cancellationTokenSource.Dispose();
        }

        public ValueTask DisposeAsync()
        {
            Dispose();

            return ValueTask.CompletedTask;
        }
    }
}