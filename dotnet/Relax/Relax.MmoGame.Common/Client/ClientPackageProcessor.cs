using System;

namespace Relax.MmoGame.Common
{
    public interface IServerService
    {
        void AddPackage(byte[] bytes);

        PlayerPosition Connect();

        void Move(MovePackage package);

        void Disconnect();
    }

    public class ClientPackageProcessor
    {
        private readonly IServerService _serverService;

        public ClientPackageProcessor(IServerService serverService)
        {
            _serverService = serverService;
        }

        public ClientCommand ReceiveProcess(Span<byte> span)
        {
            var command = (ClientCommand) span[0];

            switch (command)
            {
                case ClientCommand.Connect:
                    _serverService.Connect();
                    break;
                case ClientCommand.Disconnect:
                    _serverService.Disconnect();
                    break;
                case ClientCommand.Reconnect:
                    break;
                case ClientCommand.Move:
                    _serverService.Move(MovePackage.Deserialize(span));
                    break;
                case ClientCommand.None:
                    break;
                default:
                    throw new ArgumentOutOfRangeException();
            }

            return command;
        }
    }

    #region Packages

    public struct DisconnectPackage : IPackage<ClientCommand>
    {
        public ClientCommand Command { get; private init; }

        public int Size => 1;

        public byte[] Serialize(Span<byte> buffer)
        {
            buffer[0] = (byte) ClientCommand.Disconnect;

            return buffer[..Size].ToArray();
        }

        public static DisconnectPackage Deserialize(Span<byte> bytes)
        {
            return new()
            {
                Command = ClientCommand.Disconnect,
            };
        }
    }

    public struct MovePackage : IPackage<ClientCommand>
    {
        public ClientCommand Command { get; private init; }

        public byte PlayerId { get; init; }

        public MoveDirection Direction { get; init; }

        public int Size => 3;

        public byte[] Serialize(Span<byte> buffer)
        {
            buffer[0] = (byte) ClientCommand.Move;
            buffer[1] = PlayerId;
            buffer[2] = (byte) Direction;

            return buffer[..Size].ToArray();
        }

        public static MovePackage Deserialize(Span<byte> bytes)
        {
            return new()
            {
                Command = ClientCommand.Move,
                PlayerId = bytes[1],
                Direction = (MoveDirection) bytes[2]
            };
        }
    }

    #endregion
}