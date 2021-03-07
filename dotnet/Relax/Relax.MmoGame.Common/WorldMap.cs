using System;
using System.Collections.Generic;
using System.Linq;

namespace Relax.MmoGame.Common
{
    public class WorldMap
    {
        private readonly int _size = 10;

        public WorldMap()
        {
        }

        public void Draw(params PlayerPosition[] players)
        {
            Array.Sort(players, new PeopleComparer());

            DrawLine();
            for (int y = 0; y < _size; y++)
            {
                for (int x = 0; x < _size; x++)
                {
                    const string cell = "|_";

                    var player = players.FirstOrDefault(position => position.X == x && position.Y == y);

                    if (player != null)
                    {
                        
                    }
                    
                    Console.Write(player != null ? $"|{player.Id}" : cell);
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