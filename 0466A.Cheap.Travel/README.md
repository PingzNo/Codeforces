# 1B. Cheap Travel

## Problem

Ann has recently started commuting by subway. We know that a one ride subway ticket costs $a$ rubles. Besides, Ann found out that she can buy a special ticket for $m$ rides (she can buy it several times). It costs $b$ rubles. Ann did the math; she will need to use subway $n$ times. Help Ann, tell her what is the minimum sum of money she will have to spend to make $n$ rides?

[Link to CodeForces](https://codeforces.com/problemset/problem/466/A)

## Input

The single line contains four space-separated integers $n$, $m$, $a$, $b$ ($1 \leq n, m, a, b \leq 1000$) — the number of rides Ann has planned, the number of rides covered by the $m$ ride ticket, the price of a one ride ticket and the price of an $m$ ride ticket.

## Output

Print a single integer — the minimum sum in rubles that Ann will need to spend.

## Examples

### Input

```
6 2 1 2
```

### Output

```
6
```

### Input

```
5 2 2 3
```

### Output

```
8
```

## Note

In the first sample one of the optimal solutions is: each time buy a one ride ticket. There are other optimal solutions. For example, buy three $m$ ride tickets.

## Solutions

### Constraints

  - Time limit per test: 1 second
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

### Go 1.18.3

|  Problem  |    Lang   |  Verdict |  Time  |  Memory  |
|:---------:|:---------:|:--------:|:------:|:--------:|
| 466A - 9  |    Go     | Accepted |  31 ms |   12 KB  |

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
	var n, m, a, b int
	Fscanf("%d %d %d %d\n", &n, &m, &a, &b)

	aCost := n * a
	bCost := n/m*b + int(math.Min(float64(n%m*a), float64(b)))
	Fprintf("%d\n", int(math.Min(float64(aCost), float64(bCost))))
}

func main() {
	Solution()
	writer.Flush()
}
```
