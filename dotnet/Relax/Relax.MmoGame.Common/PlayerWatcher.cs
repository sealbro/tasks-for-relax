using System;
using System.Threading;

namespace Relax.MmoGame.Common
{
    public class PlayerWatcher
    {
        private readonly byte _sizeX;
        private readonly byte _sizeY;
        private readonly Random _random = new(DateTime.Now.Millisecond);
        private volatile int _playerCounter;

        public byte WorldSize { get; }

        public PlayerWatcher(byte size)
        {
            WorldSize = _sizeX = _sizeY = size;
        }

        public PlayerPosition RegisterPlayer()
        {
            var x = (byte) _random.Next(0, _sizeX - 1);
            var y = (byte) _random.Next(0, _sizeY - 1);

            Interlocked.Increment(ref _playerCounter);
            
            return new PlayerPosition((byte)_playerCounter, x, y);
        }

        public void MovePlayer(PlayerPosition player, MoveDirection direction)
        {
            switch (direction)
            {
                case MoveDirection.Left:
                    if (player.Y != 0)
                    {
                        player.Y -= 1;
                    }

                    break;
                case MoveDirection.Right:
                    if (player.Y != _sizeY - 1)
                    {
                        player.Y += 1;
                    }

                    break;
                case MoveDirection.Up:
                    if (player.X != 0)
                    {
                        player.X -= 1;
                    }

                    break;
                case MoveDirection.Down:
                    if (player.X != _sizeX - 1)
                    {
                        player.X += 1;
                    }

                    break;
                case MoveDirection.None:
                    break;
            }
        }
    }

    public class PlayerPosition : IComparable<PlayerPosition>
    {
        public byte PlayerId { get; }

        public byte X { get; set; }

        public byte Y { get; set; }

        public PlayerPosition(byte playerId, byte x, byte y)
        {
            PlayerId = playerId;
            X = x;
            Y = y;
        }

        public int CompareTo(PlayerPosition other)
        {
            var xComparison = X.CompareTo(other.X);
            return xComparison != 0 ? xComparison : Y.CompareTo(other.Y);
        }
    }

    public enum MoveDirection : byte
    {
        None,
        Left,
        Right,
        Up,
        Down,

        Jump
    }

    public enum ClientCommand : byte
    {
        None,
        Connect,
        Disconnect,
        Reconnect,
        Move
    }

    public enum ServerCommand : byte
    {
        None,
        Registered,
        Exit,
        Draw
    }
}