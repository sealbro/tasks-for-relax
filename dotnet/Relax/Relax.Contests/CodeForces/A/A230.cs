using System;

namespace Relax.Contests.CodeForces.A
{
    /// <summary>
    /// https://codeforces.com/problemset/problem/230/A
    /// </summary>
    public static class A230
    {
        public static void MainX(string[] args)
        {
            var inputs = Array.ConvertAll(Console.ReadLine().Split(' '), Convert.ToUInt16);
            int s = inputs[0];
            var n = inputs[1];

            ushort[] xA = new ushort[n];
            ushort[] yA = new ushort[n];

            for (ushort i = 0; i < n; i++)
            {
                inputs = Array.ConvertAll(Console.ReadLine().Split(' '), Convert.ToUInt16);
                xA[i] = inputs[0];
                yA[i] = inputs[1];
            }

            ushort count = 0;
            for (ushort i = 0; i < n; i++)
            for (ushort j = 0; j < n; j++)
            {
                if (xA[j] == 0) continue;

                if (s > xA[j])
                {
                    s += yA[j];
                    count++;
                    xA[j] = 0;
                }

                if (count == n)
                {
                    break;
                }
            }

            Console.WriteLine(count == n ? "YES" : "NO");
        }
    }
}