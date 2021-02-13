using System;

namespace Relax.Contests.CodeForces.A
{
    /// <summary>
    /// https://codeforces.com/problemset/problem/71/A
    /// </summary>
    public static class A71
    {
        public static void MainX(string[] args)
        {
            int count = Convert.ToSByte(Console.ReadLine());

            for (var i = 0; i < count; i++)
            {
                var input = Console.ReadLine();
                var length = input.Length;
                Console.WriteLine(length > 10 ? $"{input[0]}{length - 2}{input[length - 1]}" : input);
            }
        }
    }
}