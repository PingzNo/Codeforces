# 266A. Stones on the Table

## Problem

There are $n$ stones on the table in a row, each of them can be red, green or blue. Count the minimum number of stones to take from the table so that any two neighboring stones had different colors. Stones in a row are considered neighboring if there are no other stones between them.

[Link to CodeForces](https://codeforces.com/problemset/problem/266/A)

## Input

The first line contains integer $n$ ($1 \leq n \leq 50$) — the number of stones on the table.

The next line contains string s, which represents the colors of the stones. We'll consider the stones in the row numbered from $1$ to $n$ from left to right. Then the $i$-th character s equals "R", if the $i$-th stone is red, "G", if it's green and "B", if it's blue.

## Output

Print a single integer — the answer to the problem.

## Examples

### Input

```
3
RRG
```

### Output

```
1
```

### Input

```
5
RRRRR
```

### Output

```
4
```

### Input

```
4
BRBG
```

### Output

```
0
```

## Solutions

### Constraints

  - Time limit per test: 2 seconds
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

### GNU C++17 7.3.0

| Problem  |    Lang   |  Verdict | Time  | Memory |
|:--------:|:---------:|:--------:|:-----:|:------:|
| 266A - 8 |     Go    | Accepted | 62 ms |  64 KB |

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
	var stoneCount int
	var stones string
	Fscanf("%d\n%s\n", &stoneCount, &stones)

	removedCount := 0
	stone := stones[0]
	for stoneIndex := 1; stoneIndex < stoneCount; stoneIndex++ {
		if stones[stoneIndex] == stone {
			removedCount++
		} else {
			stone = stones[stoneIndex]
		}
	}

	Fprintf("%d\n", removedCount)
}

func main() {
	Solution()
	writer.Flush()
}
```
