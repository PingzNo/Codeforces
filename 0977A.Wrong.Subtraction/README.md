# 977A. Wrong Subtraction

## Problem

Little girl Tanya is learning how to decrease a number by one, but she does it wrong with a number consisting of two or more digits. Tanya subtracts one from a number by the following algorithm:

    if the last digit of the number is non-zero, she decreases the number by one;
    if the last digit of the number is zero, she divides the number by 10 (i.e. removes the last digit).

You are given an integer number $n$. Tanya will subtract one from it $k$ times. Your task is to print the result after all $k$ subtractions.

It is guaranteed that the result will be positive integer number.

[Link to CodeForces](https://codeforces.com/problemset/problem/977/A)

## Input

The first line of the input contains two integer numbers $n$ and $k$ ($2 \leq n \leq 10^9$, $1 \leq k \leq 50$) — the number from which Tanya will subtract and the number of subtractions correspondingly.

## Output

Print one integer number — the result of the decreasing $n$ by one $k$ times.

It is guaranteed that the result will be positive integer number.

## Examples

### Input

```
512 4
```

### Output

```
50
```

### Input

```
1000000000 9
```

### Output

```
1
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
| 977A - 20 |    Go     | Accepted | 30 ms |  12 KB  |

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
	var number, times int
	Fscanf("%d %d\n", &number, &times)

	remainder := number % 10
	for number > 0 && times > 0 {
		if remainder == 0 {
			times--
			number /= 10

		} else if remainder <= times {
			times -= remainder
			number -= remainder
		} else {
			number -= times
			times = 0
		}

		remainder = number % 10
	}

	Fprintf("%d\n", number)
}

func main() {
	Solution()
	writer.Flush()
}
```
