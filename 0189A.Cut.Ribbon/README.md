# 189A. Cut Ribbon

## Problem

Polycarpus has a ribbon, its length is $n$. He wants to cut the ribbon in a way that fulfils the following two conditions:

  - After the cutting each ribbon piece should have length $a$, $b$ or $c$.
  - After the cutting the number of ribbon pieces should be maximum.

Help Polycarpus and find the number of ribbon pieces after the required cutting.

[Link to CodeForces](https://codeforces.com/problemset/problem/189/A)

## Input

The first line contains four space-separated integers $n$, $a$, $b$ and $c$ ($1 \leq n$, $a$, $b$, $c \leq 4000$) — the length of the original ribbon and the acceptable lengths of
the ribbon pieces after the cutting, correspondingly. The numbers $a$, $b$ and $c$ can coincide.

## Output

Print a single number — the maximum possible number of ribbon pieces. It is guaranteed that at least one correct ribbon cutting exists.

## Examples

### Input

```
5 5 3 2
```

### Output

```
2
```

### Input

```
7 5 5 2
```

### Output

```
2
```

## Note

In the first example Polycarpus can cut the ribbon in such way: the first piece has length 2, the second piece has length 3.

In the second example Polycarpus can cut the ribbon in such way: the first piece has length 5, the second piece has length 2.

## Solutions

### Constraints

  - Time limit per test: 1 second
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

### Go 1.18.3

|  Problem  |    Lang   |  Verdict |  Time  |  Memory  |
|:---------:|:---------:|:--------:|:------:|:--------:|
| 189A - 39 |    Go     | Accepted |  31 ms |  76  KB  |

[Link to source code](solution.go)

```go
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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
	var ribbonLen int
	cutLen := make([]int, 3)
	Fscanf("%d %d %d %d\n", &ribbonLen, &cutLen[0], &cutLen[1], &cutLen[2])

	sort.Ints(cutLen)

	cutNumber := len(cutLen)
	if cutLen[2]%cutLen[1] == 0 || cutLen[2]%cutLen[0] == 0 || cutLen[2]%(cutLen[0]+cutLen[1]) == 0 || cutLen[2] > ribbonLen {
		cutNumber = 2
	}

	if cutLen[1]%cutLen[0] == 0 || cutLen[1] > ribbonLen {
		if cutNumber == 3 {
			cutLen[1] = cutLen[2]
			cutNumber = 2
		} else {
			Fprintf("%d\n", ribbonLen/cutLen[0])
			return
		}
	}

	table := make([]int, ribbonLen+1)
	for cutIndex := 0; cutIndex < cutNumber; cutIndex++ {
		table[cutLen[cutIndex]] = 1
	}

	for cutIndex := 0; cutIndex < cutNumber; cutIndex++ {
		for j := cutLen[cutIndex]; j <= ribbonLen; j++ {
			if table[j-cutLen[cutIndex]] > 0 {
				table[j] = int(math.Max(float64(table[j-cutLen[cutIndex]]+1), float64(table[j])))
			}
		}
	}

	Fprintf("%d\n", table[ribbonLen])
}

func main() {
	Solution()
	writer.Flush()
}
```
