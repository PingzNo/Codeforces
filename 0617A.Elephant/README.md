# 617A. Elephant

## Problem

An elephant decided to visit his friend. It turned out that the elephant's house is located at point 0 and his friend's house is located at point $x$($x > 0$) of the coordinate line. In one step the elephant can move 1, 2, 3, 4 or 5 positions forward. Determine, what is the minimum number of steps he need to make in order to get to his friend's house.

[Link to CodeForces](https://codeforces.com/problemset/problem/617/A)

## Input

The input contains three positive integer numbers in the first line: $n$, $m$ and $a$ ($1 \leq  n$, $m$, $a \leq 109$).

## Output

The first line of the input contains an integer $x$ ($1 \leq x \leq 1000000$) — The coordinate of the friend's house.

## Examples

### Input

```
5
```

### Output

```
1
```

### Input

```
12
```

### Output

```
3
```

## Solutions

### Constraints

  - Time limit per test: 1 second
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

### Go 1.17.5

| Problem  |    Lang   |  Verdict | Time  | Memory |
|:--------:|:---------:|:--------:|:-----:|:------:|
| 617A - 11|     Go    | Accepted | 31 ms | 64 KB  |

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
	var pos int
	Fscanf("%d\n", &pos)

	stepCount := pos / 5
	if pos%5 != 0 {
		stepCount++
	}

	Fprintf("%d\n", stepCount)
}

func main() {
	Solution()
	writer.Flush()
}
```
