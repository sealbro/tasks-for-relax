using System;
using Relax.Contests.LeetCode.Easy;
using Relax.Contests.Tests.Core;
using Xunit;
using Xunit.Abstractions;

namespace Relax.Contests.Tests.LeetCode;

public class EasyTests: BaseUnitTest
{
    public EasyTests(ITestOutputHelper output) : base(output)
    {
    }

    [Fact]
    public void MergeTwoSortedListsTests()
    {
        MergeTwoSortedLists.MainX(Array.Empty<string>());
    }
}