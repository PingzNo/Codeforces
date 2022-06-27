# 996A. Hit the Lottery

## Problem

Allen has a LOT of money. He has n dollars in the bank. For security reasons, he wants to withdraw it in cash (we will not disclose the reasons here). The denominations for dollar bills are 1, 5, 10, 20, 100. What is the minimum number of bills Allen could receive after withdrawing his entire balance?

[Link to CodeForces](https://codeforces.com/problemset/problem/996/A)

## Input

The first and only line of input contains a single integer $n$ ($1 \leq n \leq 109$).

## Output

Output the minimum number of bills that Allen could receive.


## Examples

### Input

```
125
```

### Output

```
3
```

### Input

```
43
```

### Output

```
5
```

### Input

```
1000000000
```

### Output

```
10000000
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
| 996A - 13 |    Go     | Accepted | 31 ms | 64 KB  |

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
	var money int
	Fscanf("%d\n", &money)

	dividers := [5]int{100, 20, 10, 5, 1}

	billCountSum := 0
	for _, divider := range dividers {
		billCount := money / divider
		billCountSum += billCount
		money -= billCount * divider

	}

	Fprintf("%d\n", billCountSum)
}

func main() {
	Solution()
	writer.Flush()
}
```
