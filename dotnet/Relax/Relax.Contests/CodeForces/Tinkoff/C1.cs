using System;

namespace Relax.Contests.CodeForces.Tinkoff
{
    /// <summary>
    /// https://codeforces.com/group/RVe1s8wFr4/contest/319339/problem/C
    /// </summary>
    public static class C1
    {
        public static void MainX(string[] args)
        {
            var input = Console.ReadLine();

            var buffer = new char[input.Length];

            var idx = 0;
            var last = buffer[0]; 

            var isValid = true;

            for (int i = 0; i < input.Length; i++)
            {
                var current = input[i];

                // [] 91 93
                // () 40 41
                // {} 123 125

                var abs = Math.Abs(current - last);
                if ((abs == 0 || abs > 2) && (idx < 0 || idx > 0) && current != '[' && current != '(' && current != '{')
                {
                    isValid = false;
                    break;
                }

                if (abs > 0 && abs <= 2)
                {
                    idx--;
                    last = idx > 0 ? buffer[idx - 1] : '\0';
                }
                else
                {
                    last = buffer[idx++] = current;
                }
            }

            Console.WriteLine(isValid && idx == 0 ? "yes" : "no");
        }
    }
}