# 405A. Gravity Flip

## Problem

Little Chris is bored during his physics lessons (too easy), so he has built a toy box to keep himself occupied. The box is special, since it has the ability to change gravity.

There are n columns of toy cubes in the box arranged in a line. The i-th column contains ai cubes. At first, the gravity in the box is pulling the cubes
downwards. When Chris switches the gravity, it begins to pull all the cubes to the right side of the box. The figure shows the initial and final configurations
of the cubes in the box: the cubes that have changed their position are highlighted with orange.

![Illustration](illustration.png)

Given the initial configuration of the toy cubes in the box, find the amounts of cubes in each of the n columns after the gravity switch!

[Link to CodeForces](https://codeforces.com/problemset/problem/405/A)

## Input

The first line of input contains an integer $n$ ($1 \leq n \leq 100$), the number of the columns in the box. The next line contains $n$ space-separated integer numbers. The $i$-th number $a_i$ ($1 \leq a_i \leq 100$) denotes the number of cubes in the $i$-th column.

## Output

Output $n$ integer numbers separated by spaces, where the $i$-th number is the amount of cubes in the $i$-th column after the gravity switch.

## Examples

### Input

```
4
3 2 1 2
```

### Output

```
1 2 2 3
```

### Input

```
3
2 3 8
```

### Output

```
2 3 8
```

## Note

The first example case is shown on the figure. The top cube of the first column falls to the top of the last column; the top cube of the second column falls to the top of the third column; the middle cube of the first column falls to the top of the second column.

In the second example case the gravity switch does not change the heights of the columns.

## Solutions

### Constraints

  - Time limit per test: 1 second
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

### Go 1.17.5

|  Problem  |    Lang   |  Verdict | Time  | Memory |
|:---------:|:---------:|:--------:|:-----:|:------:|
| 405A - 19 |     Go    | Accepted | 31 ms | 60 KB  |

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
	var colCount int
	Fscanf("%d\n", &colCount)

	rowCounters := [100]int{}

	var rowCount int
	for colIndex := 0; colIndex < colCount; colIndex++ {
		Fscanf("%d", &rowCount)
		for rowIndex := 0; rowIndex < rowCount; rowIndex++ {
			rowCounters[rowIndex]++
		}
	}

	for colIndex := 0; colIndex < colCount; colIndex++ {
		rowCount = 0
		for rowIndex := 0; rowIndex < len(rowCounters); rowIndex++ {
			if rowCounters[rowIndex] >= colCount-colIndex {
				rowCount++
			}

			if rowCounters[rowIndex] == 0 {
				break
			}
		}

		if colIndex > 0 {
			Fprintf(" ")
		}

		Fprintf("%d", rowCount)
	}

	Fprintf("\n")
}

func main() {
	Solution()
	writer.Flush()
}
```
