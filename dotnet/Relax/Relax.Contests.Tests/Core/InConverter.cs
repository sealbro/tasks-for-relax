using System.Collections.Generic;
using System.IO;

namespace Relax.Contests.Tests.Core
{
    internal class InConverter : TextReader
    {
        private readonly Queue<string> _inputs;

        public InConverter(params string[] inputs)
        {
            _inputs = new Queue<string>(inputs);
        }

        public override string ReadLine()
        {
            return _inputs.Dequeue();
        }
    }
}