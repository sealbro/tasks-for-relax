using System;
using System.Collections.Concurrent;
using System.Net.Sockets;
using System.Threading;
using System.Threading.Tasks;
using Relax.MmoGame.Common;

namespace Relax.MmoGame.Server
{
    public class PlayerProcessor : IDisposable, IAsyncDisposable, IServerService
    {
        private readonly byte[] _readBuffer = new byte[Consts.MaxServerRead];
        private readonly byte[] _writeBuffer = new byte[Consts.MaxServerWrite];

        #region Player

        private readonly PlayerWatcher _playerWatcher;
        private readonly ClientPackageProcessor _clientPackageProcessor;

        private readonly ConcurrentQueue<byte[]> _packagesToSend = new();
        private TcpClient _client;
        private NetworkStream _stream;
        private CancellationTokenSource _clientCancellationTokenSource;

        private readonly TimeSpan DisconnectTimeout = TimeSpan.FromSeconds(15);

        private DateTime lastReceive;

        #endregion

        public bool CanRead => Connected && _stream.CanRead && _stream.DataAvailable;

        public bool CanWrite => Connected && _stream.CanWrite;

        public bool Connected { get; private set; }

        public PlayerPosition Player { get; private set; }

        public PlayerProcessor(PlayerWatcher playerWatcher)
        {
            _playerWatcher = playerWatcher;

            _clientPackageProcessor = new ClientPackageProcessor(this);
        }

        public void Run(TcpClient client, CancellationToken cancellationToken)
        {
            _client = client;
            _stream = client.GetStream();
            _clientCancellationTokenSource = CancellationTokenSource.CreateLinkedTokenSource(cancellationToken);

            Player = Connect();
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
            lastReceive = DateTime.Now;

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

        public async Task ReceivePackages()
        {
            if (!CanRead)
            {
                return;
            }
            
            if (DateTime.Now - lastReceive > DisconnectTimeout)
            {
                Connected = false;
            }

            var cancellationToken = _clientCancellationTokenSource.Token;

            await _stream.ReadAsync(_readBuffer, cancellationToken);
            var command = _clientPackageProcessor.ReceiveProcess(_readBuffer);

            Console.WriteLine(DateTime.Now.ToString("hh:mm:ss.fff") + $": {Player.PlayerId} => {command}");

            lastReceive = DateTime.Now;
        }

        public async Task SendPackages()
        {
            if (!CanWrite)
            {
                return;
            }
            
            while (_packagesToSend.TryDequeue(out var bytes))
            {
                await _stream.WriteAsync(bytes, _clientCancellationTokenSource.Token);
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