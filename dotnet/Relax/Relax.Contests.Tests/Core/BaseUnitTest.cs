using System;
using System.Collections.Generic;
using System.Linq;
using Xunit.Abstractions;

namespace Relax.Contests.Tests.Core
{
    public class BaseUnitTest
    {
        protected readonly string[] Args = { };
        private readonly Queue<string> _outputs;
        private readonly string _separator;

        protected BaseUnitTest(ITestOutputHelper output)
        {
            // _separator = OperatingSystem.IsWindows() ? "\r\n" : "\n";
            _separator = "\n";
            _outputs = new Queue<string>();
            Console.SetOut(new OutConverter(output, _outputs));
        }

        protected string GetWriteLineMessage()
        {
            return _outputs.Dequeue();
        }

        protected string[] GetStrings(params object[] inputs)
        {
            return GetStringsInner(inputs).ToArray();
        }

        private IEnumerable<string> GetStringsInner(params object[] inputs)
        {
            foreach (var input in inputs)
            {
                if (input is string s)
                {
                    var variables = s.Split(_separator);
                    foreach (var variable in variables)
                    {
                        yield return variable;
                    }
                }
                else
                {
                    yield return input.ToString();
                }
            }
        }

        protected void SetupInputs(params string[] inputs)
        {
            Console.SetIn(new InConverter(inputs));
        }
    }
}