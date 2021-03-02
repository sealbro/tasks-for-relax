using System;
using System.Collections.Generic;
using Relax.Dotnet.System.Collections;

namespace Relax.Dotnet
{
    class Program
    {
        static void Main(string[] args)
        {
            var size = 6;
            var arr = EnumeratorExtensions.GenerateGuids().Take(size).ToArray();
            for (int i = 0; i < size; i++)
            {
                Console.WriteLine($"{i} - {arr[i]}");
            }
        }
    }
}