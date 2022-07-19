# 214A. System of Equations

## Problem

Furik loves math lessons very much, so he doesn't attend them, unlike Rubik. But now Furik wants to get a good mark for math. For that Ms. Ivanova, his math teacher, gave him a new task. Furik solved the task immediately. Can you?

You are given a system of equations:

$$
\begin{cases}
    a^2 + b = n\\
    a + b^2 = m
\end{cases}
$$

You should count, how many there are pairs of integers ($a$,$b$) ($0 \leq a$, $b$) which satisfy the system.

[Link to CodeForces](https://codeforces.com/problemset/problem/214/A)

## Input

A single line contains two integers $n$, $m$ ($1 \leq n$, $m \leq 1000$) — the parameters of the system. The numbers on the line are separated by a space.

## Output

On a single line print the answer to the problem.

## Examples

### Input

```
9 3
```

### Output

```
1
```

### Input

```
14 28
```

### Output

```
1
```

### Input

```
4 20
```

### Output

```
0
```

## Note

In the first sample the suitable pair is integers (3, 0). In the second sample the suitable pair is integers (3, 5). In the third sample there is no suitable pair.

## Solutions

### Constraints

  - Time limit per test: 2 seconds
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

### Go 1.18.3

|  Problem  |    Lang   |  Verdict |  Time  |  Memory  |
|:---------:|:---------:|:--------:|:------:|:--------:|
| 214A - 30 |    Go     | Accepted |  62 ms |   56 KB  |

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
	var n, m int
	Fscanf("%d %d\n", &n, &m)

	pairCount := 0
	for a := 0; a <= n; a++ {
		b := n - a*a
		if b < 0 {
			continue
		}

		if a+b*b == m {
			pairCount++
		}
	}

	Fprintf("%d\n", pairCount)
}

func main() {
	Solution()
	writer.Flush()
}
```
