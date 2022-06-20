# 1328A. Divisibility Problem

## Problem

You are given two positive integers $a$ and $b$. In one move you can increase a by 1 (replace a with $a+1$). Your task is to find the minimum number of moves you need to do in order to make a divisible by $b$. It is possible, that you have to make 0 moves, as $a$ is already divisible by $b$. You have to answer $t$ independent test cases.

[Link to CodeForces](https://codeforces.com/problemset/problem/1328/A)

## Input

The first line of the input contains one integer $t$ ($1 \leq t \leq 10^4$) — the number of test cases. Then $t$ test cases follow.

The only line of the test case contains two integers $a$ and $b$ ($1 \leq a, b \leq 10^9$).

## Output

For each test case print the answer — the minimum number of moves you need to do in order to make $a$ divisible by $b$.

## Examples

### Input

```
5
10 4
13 9
100 13
123 456
92 46
```

### Output

```
2
5
4
333
0
```

## Solutions

### Constraints

  - Time limit per test: 1 second
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

### GNU C++17 7.3.0

|   Problem  |    Lang   |  Verdict | Time  | Memory |
|:----------:|:---------:|:--------:|:-----:|:------:|
| 1328A - 12 |     Go    | Accepted | 31 ms | 312 KB |

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

	var a, b int
	for caseIndex := 0; caseIndex < caseCount; caseIndex++ {
		Fscanf("%d %d\n", &a, &b)
		if a%b == 0 {
			Fprintf("0\n")
		} else {
			Fprintf("%d\n", b-a%b)
		}
	}
}

func main() {
	Solution()
	writer.Flush()
}
```
