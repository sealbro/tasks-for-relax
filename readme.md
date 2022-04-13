# Tasks for relax

[![LeetCode user sealbro](https://img.shields.io/badge/dynamic/json?style=for-the-badge&labelColor=black&color=%23ffa116&label=Solved&query=solvedOverTotal&url=https%3A%2F%2Fleetcode-badge.vercel.app%2Fapi%2Fusers%2Fsealbro&logo=leetcode&logoColor=yellow)](https://leetcode.com/sealbro/)

## Additional information

- [ANSI numbers](http://www.alanwood.net/demos/ansi.html)
  - A-Z = 65-90
  - a-z = 97-122 (32 diff)
- [dotnet](https://docs.microsoft.com/en-us/dotnet/csharp/language-reference/builtin-types/integral-numeric-types) or [golang](https://www.callicoder.com/golang-basic-types-operators-type-conversion/) types suites for task constraints
  - 0 < `ushort` | `uint16` < 10^4 `[65,535]`
  - -10^4 < `short` | `int16` < 10^4 `[-32,768 to 32,767]`
  - 0 < `uint` | `uint32` < 10^9 `[4,294,967,295]`
  - -10^9 < `int` | `int32` < 10^9 `[-2,147,483,648 to 2,147,483,647]`

## Algorithms summary

- [Lexicographic Order](https://thecodingtrain.com/CodingChallenges/035.2-tsp.html)
  - Find the largest x such that P[x]<P[x+1].
  - (If there is no such x, P is the last permutation.)
  - Find the largest y such that P[x]<P[y].
  - Swap P[x] and P[y].
  - Reverse P[x+1 .. n].
- [The Combinatorial Derivation of the Fibonacci Sequence](http://mathonline.wikidot.com/the-combinatorial-derivation-of-the-fibonacci-sequence)
