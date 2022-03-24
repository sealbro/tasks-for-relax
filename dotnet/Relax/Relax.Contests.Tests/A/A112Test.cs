using Relax.Contests.CodeForces.A;
using Relax.Contests.Tests.Core;
using Xunit;
using Xunit.Abstractions;

namespace Relax.Contests.Tests.A
{
    public class A112Tests : BaseUnitTest
    {
        public A112Tests(ITestOutputHelper output) : base(output)
        {
        }

        [Theory]
        [InlineData(@"aaaa
aaaA", "0")]
        [InlineData(@"abs
Abz", "-1")]
        [InlineData(@"abcdefg
AbCdEfF", "1")]
        [InlineData(@"aslkjlkasdd
asdlkjdajwi", "1")]
        public void A112_Test(string input, string expected)
        {
            SetupInputs(GetStrings(input));

            A112.MainX(Args);

            var actual = GetWriteLineMessage();

            Assert.Equal(expected, actual);
        }
    }
}