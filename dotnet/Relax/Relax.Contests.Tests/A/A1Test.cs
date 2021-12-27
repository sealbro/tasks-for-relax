using Relax.Contests.CodeForces.A;
using Relax.Contests.Tests.Core;
using Xunit;
using Xunit.Abstractions;

namespace Relax.Contests.Tests.A
{
    public class A1Tests : BaseUnitTest
    {
        public A1Tests(ITestOutputHelper output) : base(output)
        {
        }

        [Theory]
        [InlineData(@"6 6 4", "4")]
        [InlineData(@"4 6 4", "2")]
        [InlineData(@"4 2 4", "1")]
        [InlineData(@"1000000000 1000000000 1", "1000000000000000000")]
        public void A1_Test(string input, string expected)
        {
            SetupInputs(GetStrings(input));

            A1.MainX(Args);

            var actual = GetWriteLineMessage();

            Assert.Equal(expected, actual);
        }
    }
}