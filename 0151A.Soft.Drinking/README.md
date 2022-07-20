# 151A. Soft Drinking

## Problem

This winter is so cold in Nvodsk! A group of n friends decided to buy $k$ bottles of a soft drink called "Take-It-Light" to warm up a bit. Each bottle has $l$ milliliters of the drink. Also they bought $c$ limes and cut each of them into $d$ slices. After that they found $p$ grams of salt.

To make a toast, each friend needs $nl$ milliliters of the drink, a slice of lime and $np$ grams of salt. The friends want to make as many toasts as they can, provided they all drink the same amount. How many toasts can each friend make?

[Link to CodeForces](https://codeforces.com/problemset/problem/151/A)

## Input

The first and only line contains positive integers $n$, $k$, $l$, $c$, $d$, $p$, $nl$, $np$, not exceeding 1000 and no less than 1. The numbers are separated by exactly one space.

## Output

Print a single integer — the number of toasts each friend can make.

## Examples

### Input 1

```
3 4 5 10 8 100 3 1
```

### Output 1

```
2
```

### Input 2

```
5 100 10 1 19 90 4 3
```

### Output 2

```
3
```

### Input 3

```
10 1000 1000 25 23 1 50 1
```

### Output 3

```
0
```

## Note

A comment to the first sample:

Overall the friends have 4 * 5 = 20 milliliters of the drink, it is enough to make 20 / 3 = 6 toasts. The limes are enough for 10 * 8 = 80 toasts and the salt is enough for 100 / 1 = 100 toasts. However, there are 3 friends in the group, so the answer is min(6, 80, 100) / 3 = 2.

## Solutions

### Constraints

  - Time limit per test: 2 seconds
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

### Python 3.8.10

|  Problem  |    Lang   |  Verdict |  Time  |  Memory  |
|:---------:|:---------:|:--------:|:------:|:--------:|
| 151A - 28 |    Go     | Accepted |  92 ms |    0 KB  |

[Link to source code](solution.py)

```python
def solution():
	n, k, l, c, d, p, nl, np = [int(x) for x in input().split()]

	toast_count = k * l // nl
	slice_count = c * d
	salt_count = p // np

	result = min((toast_count, slice_count, salt_count)) // n
	print(result)


if __name__ == "__main__":
	solution()
```

### Go 1.17.5

|  Problem  |    Lang   |  Verdict |  Time  |  Memory  |
|:---------:|:---------:|:--------:|:------:|:--------:|
| 151A - 28 |    Go     | Accepted |  62 ms |   64 KB  |

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
	var n, k, l, c, d, p, nl, np int
	Fscanf("%d %d %d %d %d %d %d %d\n", &n, &k, &l, &c, &d, &p, &nl, &np)

	toastCount := k * l / nl
	sliceCount := c * d
	saltCount := p / np

	result := toastCount
	if sliceCount < result {
		result = sliceCount
	}
	if saltCount < result {
		result = saltCount
	}

	result /= n

	Fprintf("%d\n", result)
}

func main() {
	Solution()
	writer.Flush()
}
```
