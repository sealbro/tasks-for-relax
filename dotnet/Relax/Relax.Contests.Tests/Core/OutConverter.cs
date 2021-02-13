using System.Collections.Generic;
using System.IO;
using System.Text;
using Xunit.Abstractions;

namespace Relax.Contests.Tests.Core
{
    internal class OutConverter : TextWriter
    {
        private readonly Queue<string> _outputs;
        private readonly ITestOutputHelper _output;

        public OutConverter(ITestOutputHelper output, Queue<string> outputs)
        {
            _output = output;
            _outputs = outputs;
        }

        public override Encoding Encoding => Encoding.Default;

        public override void WriteLine(string message)
        {
            WriteLine(message, System.Array.Empty<object>());
        }

        public override void WriteLine(string format, params object[] args)
        {
            _outputs.Enqueue(string.Format(format, args));
            _output.WriteLine(format, args);
        }
    }
}