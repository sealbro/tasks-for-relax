using System;

namespace Relax.Contests.Tinkoff
{
    /// <summary>
    /// https://codeforces.com/group/RVe1s8wFr4/contest/319339/problem/A
    /// </summary>
    public static class A1
    {
        public static void MainX(string[] args)
        {
            var inputs = Console.ReadLine().Split(' ');
            int a = Convert.ToInt16(inputs[0]);
            int b = Convert.ToInt16(inputs[1]);

            Console.WriteLine(a+b);
        }
    }
}