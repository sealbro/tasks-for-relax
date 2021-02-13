using System;

namespace Relax.Contests.CodeForces.A
{
    /// <summary>
    /// https://codeforces.com/problemset/problem/231/A
    /// </summary>
    public static class A231
    {
        public static void MainX(string[] args)
        {
            int count = Convert.ToUInt16(Console.ReadLine());

            var sum = 0;

            for (var i = 0; i < count; i++)
            {
                var input = Console.ReadLine();
                sum += input[0] + input[2] + input[4] > 145 ? 1 : 0;
            }

            Console.WriteLine(sum);
        }
    }
}