# 271A. Beautiful Year

## Problem

It seems like the year of 2013 came only yesterday. Do you know a curious fact? The year of 2013 is the first year after the old 1987 with only distinct digits.

Now you are suggested to solve the following problem: given a year number, find the minimum year number which is strictly larger than the given one and has only distinct digits.

[Link to CodeForces](https://codeforces.com/problemset/problem/1/A)

## Input

The single line contains integer $y$ ($1000 \leq y \leq 9000$) — the year number.

## Output

Print a single integer — the minimum year number that is strictly larger than y and all it's digits are distinct. It is guaranteed that the answer exists.

## Examples

### Input

```
1987
```

### Output

```
2013
```

### Input

```
2013
```

### Output

```
2014
```

## Constraints

  - Time limit per test: 2 seconds
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

## Solutions

### Go 1.17.5

|  Problem  |    Lang   |  Verdict | Time  | Memory |
|:---------:|:---------:|:--------:|:-----:|:------:|
| 271A - 11 |     Go    | Accepted | 62 ms | 64 KB  |

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
	var year int
	Fscanf("%d\n", &year)

	var a, b, c, d int
	for a = year / 1000; a <= 9; a++ {
		for b = 0; b <= 9; b++ {
			if b == a {
				continue
			}

			for c = 0; c <= 9; c++ {
				if c == a || c == b {
					continue
				}

				for d = 0; d <= 9; d++ {
					if d == a || d == b || d == c {
						continue
					}

					prefectYear := a*1000 + b*100 + c*10 + d
					if prefectYear > year {
						Fprintf("%d\n", prefectYear)
						return
					}
				}
			}
		}
	}
}

func main() {
	Solution()
	writer.Flush()
}
```
