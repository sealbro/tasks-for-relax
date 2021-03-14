using System;

namespace Relax.MmoGame.Common
{
    public static class BotPlayer
    {
        private static readonly Random _random = new Random(DateTime.Now.Millisecond);

        public static MoveDirection GetDirection()
        {
            return (MoveDirection)_random.Next(0, (int)MoveDirection.Jump);
        }
    }
}