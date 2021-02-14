using Relax.Contests.CodeForces.A;
using Relax.Contests.Tests.Core;
using Xunit;
using Xunit.Abstractions;

namespace Relax.Contests.Tests
{
    public class ATests : BaseUnitTest
    {
        public ATests(ITestOutputHelper output) : base(output)
        {
        }

        [Theory]
        [InlineData(1, "NO")]
        [InlineData(2, "NO")]
        [InlineData(13, "NO")]
        [InlineData(8, "YES")]
        [InlineData(40, "YES")]
        [InlineData(100, "YES")]
        public void A4_Test(sbyte input, string output)
        {
            SetupInputs(input.ToString());

            A4.MainX(Args);

            Assert.Equal(output, GetWriteLineMessage());
        }

        [Theory]
        [InlineData(4,
            @"word
localization
internationalization
pneumonoultramicroscopicsilicovolcanoconiosis",
            @"word
l10n
i18n
p43s")]
        public void A71_Test(sbyte count, string input, string output)
        {
            SetupInputs(GetStrings(count, input));

            A71.MainX(Args);

            var results = GetStrings(output);

            for (var i = 0; i < count; i++)
            {
                var expected = results[i];
                var actual = GetWriteLineMessage();

                Assert.Equal(expected, actual);
            }
        }

        [Theory]
        [InlineData(3,
            @"1 1 0
1 1 1
1 0 0", 2)]
        [InlineData(2,
            @"1 0 0
0 1 1", 1)]
        public void A231_Test(sbyte count, string input, int expected)
        {
            SetupInputs(GetStrings(count, input));

            A231.MainX(Args);

            var message = GetWriteLineMessage();
            var actual = int.Parse(message);

            Assert.Equal(expected, actual);
        }

        [Theory]
        [InlineData("2 4", "4")]
        [InlineData("3 3", "4")]
        public void A50_Test(string input, string expected)
        {
            SetupInputs(GetStrings(input));

            A50.MainX(Args);

            var actual = GetWriteLineMessage();

            Assert.Equal(expected, actual);
        }
    }
}