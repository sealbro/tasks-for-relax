using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using BenchmarkDotNet.Attributes;
using BenchmarkDotNet.Jobs;
using BenchmarkDotNet.Running;

namespace Relax.Benchmarks
{
    class Program
    {
        public static void Main(string[] args) => BenchmarkSwitcher.FromAssembly(typeof(Program).Assembly).Run(args);
    }
    
    [SimpleJob(RuntimeMoniker.NetCoreApp50)]
    [RPlotExporter]
    public class Strings
    {
        [Benchmark]
        [ArgumentsSource(nameof(Data))]
        public string StrBuilder(string a, string[] b)
        {
            var builder = new StringBuilder();
            builder.AppendFormat("[{0}:", a);

            for (int i = 0; i < b.Length; i++)
            {
                builder.Append(b[i]);
                if (i != b.Length)
                {
                    builder.Append(";");
                }
            }


            builder.Append("], ");
            return builder.ToString();
        }

        [Benchmark]
        [ArgumentsSource(nameof(Data))]
        public string StrInterpolation(string a, string[] b)
        {
            return $"[{a}:{string.Join(";", b)}], ";
        }

        public IEnumerable<object[]> Data()
        {
            yield return new object[] { new string('w', 300), new string[] {new('z', 300), new('x', 100), new('6', 100) } };
            // yield return new object[] { new string('2', 200), new string[] {new('z', 300), new('x', 100), new('6', 100) } };
            // yield return new object[] { new string('3', 100), new string[] {new('z', 300), new('x', 100), new('6', 100) } };
            // yield return new object[] { new string('q', 500), new string[] {new('z', 300), new('x', 100), new('6', 100) } };
            // yield return new object[] { new string('5', 600), new string[] {new('z', 300), new('x', 100), new('6', 100) } };
            // yield return new object[] { new string('6', 200), new string[] {new('z', 300), new('x', 100), new('6', 100) } };
        }
    }
}