# 230B. T-primes

## Problem

We know that prime numbers are positive integers that have exactly two distinct positive divisors. Similarly, we'll call a positive integer $t$ Т-prime, if $t$ has exactly three distinct positive divisors.

You are given an array of $n$ positive integers. For each of them determine whether it is Т-prime or not.

[Link to CodeForces](https://codeforces.com/problemset/problem/230/B)

## Input

The first line contains a single positive integer, $n$ ($1 \leq n \leq 105$), showing how many numbers are in the array. The next line contains n space-separated integers $x_i$ ($1 \leq x_i \leq 1012$).

Please, do not use the %lld specifier to read or write 64-bit integers in С++. It is advised to use the cin, cout streams or the %I64d specifier.

## Output

Print $n$ lines: the $i$-th line should contain "YES" (without the quotes), if number $x_i$ is Т-prime, and "NO" (without the quotes), if it isn't.

## Examples

### Input

```
3
4 5 6
```

### Output

```
YES
NO
NO
```

## Note

The given test has three numbers. The first number 4 has exactly three divisors — 1, 2 and 4, thus the answer for this number is "YES". The second number 5 has two divisors (1 and 5), and the third number 6 has four divisors (1, 2, 3, 6), hence the answer for them is "NO".

## Solutions

### Constraints

  - Time limit per test: 1 second
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

### GNU C++17 7.3.0

|  Problem  |    Lang   |  Verdict | Time   |  Memory  |
|:---------:|:---------:|:--------:|:------:|:--------:|
| 230B - 28 |     Go    | Accepted | 434 ms | 12188 KB |

[Link to source code](solution.go)

```go
package main

import (
	"bufio"
	"fmt"
	"math"
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
	count := int64(10e6)
	flags := make([]bool, count+1)
	flags[0] = true
	flags[1] = true

	upperLimit := int64(math.Sqrt(float64(count))) + 1
	for divider := int64(2); divider <= upperLimit; divider++ {
		if flags[divider] {
			continue
		}

		for x := divider + divider; x <= int64(count); x += divider {
			flags[x] = true
		}
	}

	var numberCount int
	Fscanf("%d\n", &numberCount)

	var number int64
	for numberIndex := 0; numberIndex < numberCount; numberIndex++ {
		Fscanf("%d", &number)
		root := int64(math.Sqrt(float64(number)))
		if root*root == number && !flags[root] {
			Fprintf("YES\n")
			continue
		}

		Fprintf("NO\n")
	}
}

func main() {
	Solution()
	writer.Flush()
}
```
