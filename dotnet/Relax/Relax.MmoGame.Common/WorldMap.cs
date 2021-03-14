using System;
using System.Collections.Generic;
using System.Linq;

namespace Relax.MmoGame.Common
{
    public class WorldMap
    {
        private readonly int _size;

        public WorldMap(int size)
        {
            _size = size;
        }

        public void Draw(PlayerPosition[] players)
        {
            Array.Sort(players, new PeopleComparer());

            Console.Clear();

            DrawLine();
            for (int x = 0; x < _size; x++)
            {
                for (int y = 0; y < _size; y++)
                {
                    const string cell = "|_";

                    // todo оптимизировать перебор
                    var player = players.FirstOrDefault(position => position.X == x && position.Y == y);


                    Console.Write(player != null ? $"|{player.PlayerId}" : cell);
                }

                Console.Write("|\n");
            }
        }

        private void DrawLine()
        {
            Console.WriteLine(new string('_', _size * 2 + 1));
        }
    }

    class PeopleComparer : IComparer<PlayerPosition>
    {
        public int Compare(PlayerPosition p1, PlayerPosition p2)
        {
            return p1.CompareTo(p2);
        }
    }
}