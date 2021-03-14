using System;

namespace Relax.MmoGame.Common
{
    public interface IClientService
    {
        void Registered(RegisteredPackage package);

        void Exit();

        void Draw(DrawPackage package);
    }

    public class ServerPackageProcessor
    {
        private readonly IClientService _clientService;

        public ServerPackageProcessor(IClientService clientService)
        {
            _clientService = clientService;
        }

        public ServerCommand ReceiveProcess(Span<byte> span)
        {
            var command = (ServerCommand) span[0];

            switch (command)
            {
                case ServerCommand.Registered:
                    _clientService.Registered(RegisteredPackage.Deserialize(span));
                    break;
                case ServerCommand.Exit:
                    _clientService.Exit();
                    break;
                case ServerCommand.Draw:
                    _clientService.Draw(DrawPackage.Deserialize(span));
                    break;
                case ServerCommand.None:
                    break;
                default:
                    throw new ArgumentOutOfRangeException();
            }

            return command;
        }
    }

    #region Packages

    public struct RegisteredPackage : IPackage<ServerCommand>
    {
        public ServerCommand Command { get; private init; }

        public byte PlayerId { get; init; }

        public byte WorldSize { get; init; }

        public int Size => 3;

        public byte[] Serialize(Span<byte> buffer)
        {
            buffer[0] = (byte) ServerCommand.Registered;
            buffer[1] = PlayerId;
            buffer[2] = WorldSize;

            return buffer[..Size].ToArray();
        }

        public static RegisteredPackage Deserialize(Span<byte> bytes)
        {
            return new()
            {
                Command = ServerCommand.Registered,
                PlayerId = bytes[1],
                WorldSize = bytes[2]
            };
        }
    }

    public struct ExitPackage : IPackage<ServerCommand>
    {
        public ServerCommand Command { get; private init; }

        public byte PlayerId { get; init; }

        public int Size => 2;

        public byte[] Serialize(Span<byte> buffer)
        {
            buffer[0] = (byte) ServerCommand.Exit;
            buffer[1] = PlayerId;

            return buffer[..Size].ToArray();
        }

        public static ExitPackage Deserialize(Span<byte> bytes)
        {
            return new()
            {
                Command = ServerCommand.Exit,
                PlayerId = bytes[1]
            };
        }
    }

    public struct DrawPackage : IPackage<ServerCommand>
    {
        public ServerCommand Command { get; private init; }

        public PlayerPosition[] Players { get; init; }

        public int Size => 2 + Players.Length * 3;

        public byte[] Serialize(Span<byte> buffer)
        {
            buffer[0] = (byte) ServerCommand.Draw;
            buffer[1] = (byte) Players.Length; // todo check int

            var idx = 2;
            foreach (var player in Players)
            {
                buffer[idx++] = player.PlayerId;
                buffer[idx++] = player.X;
                buffer[idx++] = player.Y;
            }

            return buffer[..Size].ToArray();
        }

        public static DrawPackage Deserialize(Span<byte> bytes)
        {
            var count = bytes[1];

            var players = new PlayerPosition[count];

            var idx = 2;
            for (var i = 0; i < count; i++)
            {
                players[i] = new PlayerPosition(bytes[idx++], bytes[idx++], bytes[idx++]);
            }

            return new()
            {
                Command = ServerCommand.Draw,
                Players = players
            };
        }
    }

    #endregion
}