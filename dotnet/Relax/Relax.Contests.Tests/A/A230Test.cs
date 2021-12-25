using Relax.Contests.CodeForces.A;
using Relax.Contests.Tests.Core;
using Xunit;
using Xunit.Abstractions;

namespace Relax.Contests.Tests.A
{
    public class A230Tests : BaseUnitTest
    {
        public A230Tests(ITestOutputHelper output) : base(output)
        {
        }

        [Theory]
        [InlineData(@"2 2
1 99
100 0", "YES")]
        [InlineData(@"10 1
100 100", "NO")]
        [InlineData(@"999 2
1010 10
67 89", "YES")]
        public void A230_Test(string input, string expected)
        {
            SetupInputs(GetStrings(input));

            A230.MainX(Args);

            var actual = GetWriteLineMessage();

            Assert.Equal(expected, actual);
        }
    }
}