using System;

namespace Relax.MmoGame.Common
{
    public class PlayerWatcher
    {
        public RegisteredPlayer RegisterPlayer()
        {
            // инкрементировать ИД
            //

            return new RegisteredPlayer();
        }

        public void MovePlayer(int playerId, MoveDirection direction)
        {
        }
    }

    public class RegisteredPlayer
    {
        public int Id { get; set; }
        public int X { get; set; }
        public int Y { get; set; }
        public MoveDirection Direction { get; set; }
    }

    public class PlayerPosition : IComparable<PlayerPosition>
    {
        public int Id { get; set; }
        public int X { get; set; }
        public int Y { get; set; }

        public int CompareTo(PlayerPosition other)
        {
            var xComparison = X.CompareTo(other.X);
            return xComparison != 0 ? xComparison : Y.CompareTo(other.Y);
        }
    }

    public enum MoveDirection
    {
        None,
        Left,
        Right,
        Up,
        Down
    }
}