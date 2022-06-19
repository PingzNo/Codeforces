# 1335A. Candies and Two Sisters

## Problem

There are two sisters Alice and Betty. You have $n$ candies. You want to distribute these $n$ candies between two sisters in such a way that:

  - Alice will get $a$ ($a > 0$) candies;
  - Betty will get $b$ ($b > 0$) candies;
  - Each sister will get some integer number of candies;
  - Alice will get a greater amount of candies than Betty (i.e. $a > b$);
  - All the candies will be given to one of two sisters (i.e. $a + b = n$).

Your task is to calculate the number of ways to distribute exactly $n$ candies between sisters in a way described above. Candies are indistinguishable.

Formally, find the number of ways to represent $n$ as the sum of $n = a + b$, where $a$ and $b$ are positive integers and $a > b$.

You have to answer $t$ independent test cases.

[Link to CodeForces](https://codeforces.com/problemset/problem/1335/A)

## Input

The first line of the input contains one integer $t$ ($1 \leq t \leq 104$) — the number of test cases. Then $t$ test cases follow.

The only line of a test case contains one integer $n$ ($1 \leq n \leq 2 \cdot 10^9$) — the number of candies you have.

## Output

For each test case, print the answer — the number of ways to distribute exactly $n$ candies between two sisters in a way described in the problem statement. If there is no way to satisfy all the conditions, print 0.

## Examples

### Input

```
6
7
1
2
3
2000000000
763243547
```

### Output

```
3
0
0
1
999999999
381621773
```

## Note

For the test case of the example, the 3 possible ways to distribute candies are:

  - $a = 6$, $b = 1$;
  - $a = 5$, $b = 2$;
  - $a = 4$, $b = 3$.

## Solutions

### Constraints

  - Time limit per test: 1 second
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

### Go 1.17.5

|   Problem  |    Lang   |  Verdict | Time  | Memory  |
|:----------:|:---------:|:--------:|:-----:|:-------:|
| 1335A - 13 |    Go     | Accepted | 31 ms | 132 KB  |

[Link to source code](solution.go)

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Fprintf(format string, x ...interface{}) {
	fmt.Fprintf(writer, format, x...)
}

func Fscanf(format string, x ...interface{}) {
	fmt.Fscanf(reader, format, x...)
}

func Solution() {
	var caseCount int
	Fscanf("%d\n", &caseCount)

	var candyCount int
	for caseIndex := 0; caseIndex < caseCount; caseIndex++ {
		Fscanf("%d\n", &candyCount)
		if candyCount%2 == 0 {
			Fprintf("%d\n", candyCount/2-1)
		} else {
			Fprintf("%d\n", candyCount/2)
		}
	}
}

func main() {
	Solution()
	writer.Flush()
}
```
