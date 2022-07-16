# 337A. Puzzles

## Problem

The end of the school year is near and Ms. Manana, the teacher, will soon have to say goodbye to a yet another class. She decided to prepare a goodbye present for her $n$ students and give each of them a jigsaw puzzle (which, as wikipedia states, is a tiling puzzle that requires the assembly of numerous small, often oddly shaped, interlocking and tessellating pieces).

The shop assistant told the teacher that there are m puzzles in the shop, but they might differ in difficulty and size. Specifically, the first jigsaw puzzle consists of $f_1$ pieces, the second one consists of $f_2$ pieces and so on.

Ms. Manana doesn't want to upset the children, so she decided that the difference between the numbers of pieces in her presents must be as small as possible. Let $A$ be the number of pieces in the largest puzzle that the teacher buys and $B$ be the number of pieces in the smallest such puzzle. She wants to choose such n puzzles that $A - B$ is minimum possible. Help the teacher and find the least possible value of $A - B$.

[Link to CodeForces](https://codeforces.com/problemset/problem/337/A)

## Input

The first line contains space-separated integers $n$ and $m$ ($2 \leq n \leq m \leq 50$). The second line contains $m$ space-separated integers $f_1$, $f_2$, ..., $f_m$ ($4 \leq f_i \leq 1000$) — the quantities of pieces in the puzzles sold in the shop.

## Output

Print a single integer — the least possible difference the teacher can obtain.

## Examples

### Input

```
4 6
10 12 10 7 5 22
```

### Output

```
5
```

## Note

Sample 1. The class has 4 students. The shop sells 6 puzzles. If Ms. Manana buys the first four puzzles consisting of 10, 12, 10 and 7 pieces correspondingly,
then the difference between the sizes of the largest and the smallest puzzle will be equal to 5. It is impossible to obtain a smaller difference. Note that the
teacher can also buy puzzles 1, 3, 4 and 5 to obtain the difference 5.

## Solutions

### Constraints

  - Time limit per test: 1 second
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

### Go 1.17.5

| Problem  |    Lang   |  Verdict | Time  | Memory |
|:--------:|:---------:|:--------:|:-----:|:------:|
| 337A - 9 |    Go     | Accepted | 30 ms | 20 KB  |

[Link to source code](solution.go)

```go
package main

import (
	"bufio"
	"fmt"
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
	var childCount, puzzleCount int
	Fscanf("%d %d\n", &childCount, &puzzleCount)

	puzzles := make([]int, puzzleCount)
	for i := 0; i < puzzleCount; i++ {
		Fscanf("%d", &puzzles[i])
	}

	sort.Ints(puzzles)

	minDiff := 1000
	for i := 0; i+childCount-1 < puzzleCount; i++ {
		if puzzles[i+childCount-1]-puzzles[i] < minDiff {
			minDiff = puzzles[i+childCount-1] - puzzles[i]
		}
	}

	Fprintf("%d\n", minDiff)
}

func main() {
	Solution()
	writer.Flush()
}
```
