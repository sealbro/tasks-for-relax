using System;

namespace Relax.Contests.CodeForces.A
{
    /// <summary>
    /// https://codeforces.com/problemset/problem/1/A
    /// </summary>
    public static class A1
    {
        public static void MainX(string[] args)
        {
            var numbers = Array.ConvertAll(Console.ReadLine().Split(' '), Convert.ToInt32);

            var n = numbers[0];
            var m = numbers[1];
            var a = numbers[2];

            Console.WriteLine(GreatDiv(n, a) * GreatDiv(m, a));
        }

        private static long GreatDiv(int a, int b)
        {
            return a / b + (a % b > 0 ? 1 : 0);
        }
    }
}