# 546A. Soldier and Bananas

## Problem

A soldier wants to buy $w$ bananas in the shop. He has to pay $k$ dollars for the first banana, $2k$ dollars for the second one and so on (in other words, he
has to pay $i \times k$ dollars for the $i$-th banana).

He has $n$ dollars. How many dollars does he have to borrow from his friend soldier to buy w bananas?

[Link to CodeForces](https://codeforces.com/problemset/problem/546/A)

## Input

The first line contains three positive integers $k$, $n$, $w$ ($1 \leq k$, $w \leq 1000$, $0 \leq n \leq 109$), the cost of the first banana, initial number of dollars the soldier has and number of bananas he wants.

## Output

Output one integer — the amount of dollars that the soldier must borrow from his friend. If he doesn't have to borrow money, output 0.

## Examples

### Input

```
3 17 4
```

### Output

```
13
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
| 546A - 23 |    Go     | Accepted | 31 ms | 56 KB  |

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
	var k, n, w int
	Fscanf("%d %d %d\n", &k, &n, &w)

	cost := k * (1 + w) * w / 2
	if cost > n {
		Fprintf("%d\n", cost-n)
	} else {
		Fprintf("0\n")
	}
}

func main() {
	Solution()
	writer.Flush()
}
```
