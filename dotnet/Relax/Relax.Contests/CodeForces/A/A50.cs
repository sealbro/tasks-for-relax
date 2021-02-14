using System;

namespace Relax.Contests.CodeForces.A
{
    /// <summary>
    /// https://codeforces.com/problemset/problem/50/A
    /// </summary>
    public static class A50
    {
        public static void MainX(string[] args)
        {
            var inputs = Console.ReadLine().Split(" ");
            int m = Convert.ToSByte(inputs[0]);
            int n = Convert.ToSByte(inputs[1]);

            Console.WriteLine(m * n / 2);
        }
    }
}