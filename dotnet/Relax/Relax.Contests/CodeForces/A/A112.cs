using System;

namespace Relax.Contests.CodeForces.A
{
    /// <summary>
    /// https://codeforces.com/problemset/problem/112/A
    /// </summary>
    public static class A112
    {
        public static void MainX(string[] args)
        {
            var a = Console.ReadLine();
            var b = Console.ReadLine();

            int sum = 0;
            for (byte i = 0; i < a.Length; i++)
            {
                sum = (a[i] < 97 ? a[i] : a[i] - 32) - (b[i] < 97 ? b[i] : b[i] - 32);

                if (sum != 0)
                {
                    break;
                }
            }

            if (sum > 0)
            {
                sum = 1;
            }
            else if (sum < 0)
            {
                sum = -1;
            }
            else
            {
                sum = 0;
            }

            Console.WriteLine(sum);
        }
    }
}