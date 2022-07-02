# 25A. IQ test

## Problem

Bob is preparing to pass IQ test. The most frequent task in this test is to find out which one of the given n numbers differs from the others. Bob observed that one number usually differs from the others in evenness. Help Bob — to check his answers, he needs a program that among the given n numbers finds one that is different in evenness.

[Link to CodeForces](https://codeforces.com/problemset/problem/25/A)

## Input

The first line contains integer $n$ ($3 \leq n \leq 100$) — amount of numbers in the task. The second line contains n space-separated natural numbers, not exceeding 100.
It is guaranteed, that exactly one of these numbers differs from the others in evenness.

## Output

Output index of number that differs from the others in evenness. Numbers are numbered from 1 in the input order.

## Examples

### Input

```
5
2 4 7 8 10
```

### Output

```
3
```

### Input

```
4
1 2 1 1
```

### Output

```
2
```

## Solutions

### Constraints

  - Time limit per test: 2 seconds
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

### Go 1.17.5

| Problem |    Lang   |  Verdict | Time  | Memory |
|:-------:|:---------:|:--------:|:-----:|:------:|
| 25A - 7 |    Go     | Accepted | 62 ms | 64 KB  |

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
	var numberCount int
	Fscanf("%d\n", &numberCount)

	counters := [2]int{}
	indices := [2]int{}

	var number int
	for numberIndex := 0; numberIndex < numberCount; numberIndex++ {
		Fscanf("%d", &number)
		index := number % 2
		counters[index]++
		if counters[index] == 1 {
			indices[index] = numberIndex + 1
		}

		if counters[index] > 1 && counters[1-index] == 1 {
			Fprintf("%d\n", indices[1-index])
			break
		}

		if counters[index] == 1 && counters[1-index] > 1 {
			Fprintf("%d\n", indices[index])
			break
		}
	}
}

func main() {
	Solution()
	writer.Flush()
}
```
