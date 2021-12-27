using System;

namespace Relax.Contests.Tinkoff
{
    /// <summary>
    /// https://codeforces.com/group/RVe1s8wFr4/contest/319339/problem/B
    /// </summary>
    public static class B1
    {
        public static void MainX(string[] args)
        {
            var stack = new XStack();

            while (true)
            {
                var input = Console.ReadLine();

                if (input == "exit")
                {
                    Console.WriteLine("bye");
                    break;
                }

                if (input.StartsWith("push"))
                {
                    Console.WriteLine("ok");
                    stack.Push(Convert.ToInt32(input.Split(" ")[1]));
                    continue;
                }

                switch (input)
                {
                    case "clear":
                        stack.Clear();
                        Console.WriteLine("ok");
                        break;
                    case "size":
                        Console.WriteLine(stack.Size);
                        break;
                    case "pop":
                        if (stack.Size == 0) Console.WriteLine("error");
                        else Console.WriteLine(stack.Pop());
                        break;
                    case "back":
                        if (stack.Size == 0) Console.WriteLine("error");
                        else Console.WriteLine(stack.Back());
                        break;
                }
            }
        }

        public class XStack
        {
            private int[] _values = Array.Empty<int>();

            public int Size { get; private set; } = 0;

            public int Back() => _values[Size - 1];

            public void Push(int value)
            {
                CheckAndResize();

                _values[Size++] = value;
            }

            public int Pop()
            {
                return _values[--Size];
            }

            public void Clear()
            {
                Size = 0;
            }

            private void CheckAndResize()
            {
                if (Size == _values.Length)
                {
                    Array.Resize(ref _values, Size == 0 ? 2 : Size * 2);
                }
            }
        }
    }
}