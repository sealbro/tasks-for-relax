using System;
using System.Collections.Concurrent;
using System.Linq;
using System.Net.Sockets;
using System.Threading;
using System.Threading.Tasks;
using Relax.MmoGame.Common;

namespace Relax.MmoGame.Server
{
    public class PlayerProcessor : IDisposable, IAsyncDisposable, IServerService
    {
        private readonly byte[] _readBuffer = new byte[50]; // MTU 1500 bytes
        private readonly byte[] _writeBuffer = new byte[500];

        #region Player

        private readonly PlayerWatcher _playerWatcher;
        private readonly ClientPackageProcessor _clientPackageProcessor;

        private readonly ConcurrentQueue<byte[]> _packagesToSend = new();
        private TcpClient _client;
        private CancellationTokenSource _clientCancellationTokenSource;

        private TimeSpan DisconnectTimeout = TimeSpan.FromSeconds(15);

        #endregion

        public bool Connected { get; private set; }

        public PlayerPosition Player { get; }

        public PlayerProcessor(PlayerWatcher playerWatcher)
        {
            _playerWatcher = playerWatcher;

            _clientPackageProcessor = new ClientPackageProcessor(this);

            Player = Connect();
        }

        public async void Run(TcpClient client, CancellationToken cancellationToken)
        {
            _client = client;
            _clientCancellationTokenSource = CancellationTokenSource.CreateLinkedTokenSource(cancellationToken);
            var innerCancellationToken = _clientCancellationTokenSource.Token;

            await Task.Run(async () =>
            {
                await using var stream = client.GetStream();

                await Task.WhenAll(
                    ReceivePackages(stream, innerCancellationToken),
                    SendPackages(stream, innerCancellationToken)
                );

                // todo client.Dispose(); // watch cancellationToken 
            }, innerCancellationToken);
        }

        #region IServerService

        public void AddPackage(byte[] bytes)
        {
            if (!Connected) return;

            _packagesToSend.Enqueue(bytes);
        }

        public PlayerPosition Connect()
        {
            var player = _playerWatcher.RegisterPlayer();

            Connected = true;

            _packagesToSend.Enqueue(
                new RegisteredPackage
                    {PlayerId = player.PlayerId, WorldSize = _playerWatcher.WorldSize}.Serialize(_writeBuffer));

            return player;
        }

        public void Disconnect()
        {
            Connected = false;

            _packagesToSend.Enqueue(new ExitPackage {PlayerId = Player.PlayerId}.Serialize(_writeBuffer));
        }

        public void Move(MovePackage package)
        {
            if (!Connected) return;

            _playerWatcher.MovePlayer(Player, package.Direction);
        }

        #endregion

        private async Task ReceivePackages(NetworkStream stream, CancellationToken cancellationToken)
        {
            var lastReceive = DateTime.Now;

            while (!cancellationToken.IsCancellationRequested && stream.CanRead)
            {
                if (DateTime.Now - lastReceive > DisconnectTimeout)
                {
                    Connected = false;
                }

                await Task.Delay(100, cancellationToken);

                if (!stream.DataAvailable) continue;

                await stream.ReadAsync(_readBuffer, cancellationToken);

                lastReceive = DateTime.Now;

                var command = _clientPackageProcessor.ReceiveProcess(_readBuffer);

                Console.WriteLine(DateTime.Now.ToString("hh:mm:ss.fff") + $": {Player.PlayerId} => {command}");
            }
        }

        private async Task SendPackages(NetworkStream stream, CancellationToken cancellationToken)
        {
            while (!cancellationToken.IsCancellationRequested && stream.CanWrite)
            {
                while (_packagesToSend.TryDequeue(out var bytes))
                {
                    await stream.WriteAsync(bytes, cancellationToken);
                }
            }
        }

        public void Dispose()
        {
            _client?.Dispose();
            _clientCancellationTokenSource?.Dispose();
        }

        public ValueTask DisposeAsync()
        {
            Dispose();

            return ValueTask.CompletedTask;
        }
    }
}