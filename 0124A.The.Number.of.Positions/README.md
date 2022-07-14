# 124A. The number of positions

## Problem

Petr stands in line of $n$ people, but he doesn't know exactly which position he occupies. He can say that there are no less than a people standing in front of him and no more than $b$ people standing behind him. Find the number of different positions Petr can occupy.

[Link to CodeForces](https://codeforces.com/problemset/problem/124/A)

## Input

The only line contains three integers $n$, $a$ and $b$ ($0 \leq a$, $b \lt n \leq 100$).

## Output

Print the single number — the number of the sought positions.

## Examples

### Input

```
3 1 1
```

### Output

```
2
```

### Input

```
5 2 3
```

### Output

```
3
```

## Note

The possible positions in the first sample are: 2 and 3 (if we number the positions starting with 1).

In the second sample they are 3, 4 and 5.

## Solutions

### Constraints

  - Time limit per test: 0.5 second
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

### Go 1.17.5

|  Problem  |    Lang   |  Verdict |  Time  |  Memory  |
|:---------:|:---------:|:--------:|:------:|:--------:|
| 124A - 21 |    Go     | Accepted |  62 ms |   64 KB  |

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
	var peopleCount, a, b int
	Fscanf("%d %d %d\n", &peopleCount, &a, &b)

	lowerBound := a + 1
	upperBound := peopleCount
	if lowerBound > upperBound {
		Fprintf("0\n")
		return
	}

	if upperBound-lowerBound > b {
		lowerBound = peopleCount - b
	}

	Fprintf("%d\n", upperBound-lowerBound+1)
}

func main() {
	Solution()
	writer.Flush()
}
```
