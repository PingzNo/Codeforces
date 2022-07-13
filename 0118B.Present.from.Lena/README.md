# 118B. Present from Lena

## Problem

Vasya's birthday is approaching and Lena decided to sew a patterned handkerchief to him as a present. Lena chose digits from 0 to $n$ as the pattern. The digits
will form a rhombus. The largest digit $n$ should be located in the centre. The digits should decrease as they approach the edges. For example, for $n$ = 5 the
handkerchief pattern should look like that:

          0
        0 1 0
      0 1 2 1 0
    0 1 2 3 2 1 0
  0 1 2 3 4 3 2 1 0
0 1 2 3 4 5 4 3 2 1 0
  0 1 2 3 4 3 2 1 0
    0 1 2 3 2 1 0
      0 1 2 1 0
        0 1 0
          0

Your task is to determine the way the handkerchief will look like by the given $n$.

[Link to CodeForces](https://codeforces.com/problemset/problem/118/B)

## Input

The first line contains the single integer $n$ ($2 \leq n \leq 9$).

## Output

Print a picture for the given $n$. You should strictly observe the number of spaces before the first digit on each line. Every two adjacent digits in the same line should be separated by exactly one space. There should be no spaces after the last digit at the end of each line.

## Examples

### Input

```
2
```

### Output

```
    0
  0 1 0
0 1 2 1 0
  0 1 0
    0
```

### Input

```
3
```

### Output

```
      0
    0 1 0
  0 1 2 1 0
0 1 2 3 2 1 0
  0 1 2 1 0
    0 1 0
      0
```

## Solutions

### Constraints

  - Time limit per test: 2 seconds
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

### Go 1.17.5

|  Problem  |    Lang   |  Verdict |  Time  |  Memory  |
|:---------:|:---------:|:--------:|:------:|:--------:|
| 118B - 12 |    Go     | Accepted  | 60 ms |   64 KB  |

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
	var n int
	Fscanf("%d\n", &n)

	for i := 0; i <= n; i++ {
		for j := n - i; j > 0; j-- {
			Fprintf("  ")
		}

		if i != 0 {
			for j := 0; j <= i; j++ {
				Fprintf("%d ", j)
			}
		}

		for j := i - 1; j > 0; j-- {
			Fprintf("%d ", j)
		}

		Fprintf("0\n")
	}

	for i := n - 1; i >= 0; i-- {
		for j := 1; j <= n-i; j++ {
			Fprintf("  ")
		}

		if i != 0 {
			for j := 0; j <= i; j++ {
				Fprintf("%d ", j)
			}
		}

		for j := i - 1; j > 0; j-- {
			Fprintf("%d ", j)
		}

		Fprintf("0\n")
	}
}

func main() {
	Solution()
	writer.Flush()
}
```
