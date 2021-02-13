using System;

namespace Relax.Contests.CodeForces.A
{
    /// <summary>
    /// https://codeforces.com/problemset/problem/4/A
    /// </summary>
    public static class A4
    {
        public static void MainX(string[] args)
        {
            int w = Convert.ToSByte(Console.ReadLine());
            Console.WriteLine(w > 2 && w % 2 == 0  ? "YES" : "NO");
        }
    }
}