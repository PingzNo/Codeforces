# 339B. Xenia and Ringroad

## Problem

Xenia lives in a city that has $n$ houses built along the main ringroad. The ringroad houses are numbered 1 through $n$ in the clockwise order. The ringroad traffic is one way and also is clockwise.

Xenia has recently moved into the ringroad house number 1. As a result, she's got m things to do. In order to complete the $i$-th task, she needs to be in the house number ai and complete all tasks with numbers less than $i$. Initially, Xenia is in the house number 1, find the minimum time she needs to complete all her tasks if moving from a house to a neighboring one along the ringroad takes one unit of time.

[Link to CodeForces](https://codeforces.com/problemset/problem/339/B)

## Input

The first line contains two integers $n$ and $m$ ($2 \leq n \leq 10^5$, $1 \leq m \leq 10^5$). The second line contains $m$ integers $a_1$, $a_2$, ..., $a_m$ ($1 \leq a_i \leq n$). Note that Xenia can have multiple consecutive tasks in one house.

## Output

Print a single integer — the time Xenia needs to complete all tasks.

Please, do not use the %lld specifier to read or write 64-bit integers in С++. It is preferred to use the cin, cout streams or the %I64d specifier.

## Examples

### Input

```
4 3
3 2 3
```

### Output

```
6
```

### Input

```
4 3
2 3 3
```

### Output

```
2
```

## Note

In the first test example the sequence of Xenia's moves along the ringroad looks as follows: $1 \rightarrow 2 \rightarrow 3 \rightarrow 4 \rightarrow 1 \rightarrow 2 \rightarrow 3$. This is optimal sequence. So, she needs 6 time units.

## Solutions

### Constraints

  - Time limit per test: 2 seconds
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

### Go 1.17.5

|  Problem  |    Lang   |  Verdict | Time   |  Memory  |
|:---------:|:---------:|:--------:|:------:|:--------:|
| 339B - 15 |    Go     | Accepted | 156 ms |  540 KB  |

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
	var houseCount, taskCount int64
	Fscanf("%d %d\n", &houseCount, &taskCount)

	currentIndex := int64(1)

	var houseIndex, timeCost int64
	for taskIndex := int64(0); taskIndex < taskCount; taskIndex++ {
		Fscanf("%d", &houseIndex)
		if currentIndex <= houseIndex {
			timeCost += houseIndex - currentIndex
		} else {
			timeCost += houseCount - currentIndex + houseIndex
		}

		currentIndex = houseIndex
	}

	Fprintf("%d\n", timeCost)
}

func main() {
	Solution()
	writer.Flush()
}
```
