# 486A. Calculating Function

## Problem

For a positive integer $n$ let's define a function $f$:

$f(n) =  - 1 + 2 - 3 + .. + (-1)_{n}n$

Your task is to calculate $f(n)$ for a given integer $n$.

[Link to CodeForces](https://codeforces.com/problemset/problem/486/A)

## Input

The single line contains the positive integer $n$ ($1 \leq n \leq 1015$).

## Output

Print $f(n)$ in a single line.

## Examples

### Input

```
4
```

### Output

```
2
```

### Input

```
5
```

### Output

```
-3
```

## Solutions

### Constraints

  - Time limit per test: 1 second
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

### Go 1.17.5

|  Problem  |    Lang   |  Verdict | Time  | Memory |
|:---------:|:---------:|:--------:|:-----:|:------:|
| 486A - 24 |    Go     | Accepted | 15 ms | 64 KB  |

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
	var number int64
	Fscanf("%d\n", &number)

	var result int64
	if number % 2 == 0 {
		result = number/2
	} else {
		result = number/2 - number
	}

	Fprintf("%d\n", result)
}

func main() {
	Solution()
	writer.Flush()
}
```
