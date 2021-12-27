using Relax.Contests.Tests.Core;
using Relax.Contests.Tinkoff;
using Xunit;
using Xunit.Abstractions;

namespace Relax.Contests.Tests
{
    public class TinkoffTests : BaseUnitTest
    {
        public TinkoffTests(ITestOutputHelper output) : base(output)
        {
        }

        [Theory]
        [InlineData("7 8", "15")]
        [InlineData("-100 100", "0")]
        [InlineData("-7 -99", "-106")]
        public void A1_Test(string input, string output)
        {
            SetupInputs(input);

            A1.MainX(Args);

            Assert.Equal(output, GetWriteLineMessage());
        }
        
        [Theory]
        [InlineData(@"size
push 1
size
push 2
size
push 3
size
exit", @"0
ok
1
ok
2
ok
3
bye")]
        [InlineData(@"back
push 1
push 2
push 999999
size
back
pop
size
clear
pop
exit", @"error
ok
ok
ok
3
999999
999999
2
ok
error
bye")]
        public void B1_Test(string input, string output)
        {
            SetupInputs(GetStrings(input));

            B1.MainX(Args);

            var results = GetStrings(output);

            foreach (var expected in results)
            {
                var actual = GetWriteLineMessage();

                Assert.Equal(expected, actual);
            }
        }
        
        [Theory]
        [InlineData("()[]", "yes")]
        [InlineData("([)]", "no")]
        [InlineData("({[]{}})", "yes")]
        [InlineData("([})", "no")]
        [InlineData("((([[[{{}}({[]})]]])))", "yes")]
        [InlineData("((([[[{{}}({[]})]]])))((((((((((([[[{{}}({[]})]]])))(((((()))))))[][][][][[[[[[[[[[[[]]]]]]]]]]", "no")]
        [InlineData("((([[[{{}}({[]})]]])))((((((((((((((", "no")]
        [InlineData("((([[[{{}}({[]})]]])))(((((((((((((())))))))))))))))))))))))))))))", "no")]
        [InlineData("((()))))))))))))((((", "no")]
        [InlineData("()()()()(((()))(()()()()(){}{{}{][][][][}{{[[[]]]]]]]]]]]]))))))", "no")]
        [InlineData("[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[", "no")]
        [InlineData("]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]", "no")]
        public void C1_Test(string input, string output)
        {
            SetupInputs(input);

            C1.MainX(Args);

            Assert.Equal(output, GetWriteLineMessage());
        }
    }
}