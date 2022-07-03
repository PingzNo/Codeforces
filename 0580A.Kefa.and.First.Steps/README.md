# 580A. Kefa and First Steps

## Problem

Kefa decided to make some money doing business on the Internet for exactly $n$ days. He knows that on the $i$-th day ($1 \leq i \leq n$) he makes $a_i$ money. Kefa loves progress, that's why he wants to know the length of the maximum non-decreasing subsegment in sequence $a_$_i$. Let us remind you that the subsegment of the sequence is its continuous fragment. A subsegment of numbers is called non-decreasing if all numbers in it follow in the non-decreasing order.

Help Kefa cope with this task!

[Link to CodeForces](https://codeforces.com/problemset/problem/580/A)

## Input

The first line contains integer $n$ ($1 \leq n \leq 105$).

The second line contains $n$ integers $a_1$,  $a_2$,  ...,  $a_n$ ($1 \leq a_i \leq 109$).

## Output

Print a single integer — the length of the maximum non-decreasing subsegment of sequence $a$.

## Examples

### Input

```
6
2 2 1 3 4 1
```

### Output

```
3
```

### Input

```
3
2 2 9
```

### Output

```
3
```

## Note

In the first test the maximum non-decreasing subsegment is the numbers from the third to the fifth one.

In the second test the maximum non-decreasing subsegment is the numbers from the first to the third one.

## Solutions

### Constraints

  - Time limit per test: 2 seconds
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

### Go 1.17.5

|  Problem  |    Lang   |  Verdict | Time  |  Memory   |
|:---------:|:---------:|:--------:|:-----:|:---------:|
| 580A - 19 |    Go     | Accepted | 78 ms |  1568 KB  |

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
	var stepCount, prevIncome, income, beginIndex int
	Fscanf("%d\n%d", &stepCount, &prevIncome)

	maxLength := 1
	stepIndex := 1
	for stepIndex < stepCount {
		Fscanf("%d", &income)
		if income < prevIncome {
			if stepIndex-beginIndex+1 > maxLength {
				maxLength = stepIndex - beginIndex
			}

			beginIndex = stepIndex
		}

		prevIncome = income
		stepIndex++
	}

	if stepIndex-beginIndex+1 > maxLength {
		maxLength = stepIndex - beginIndex
	}

	Fprintf("%d\n", maxLength)
}

func main() {
	Solution()
	writer.Flush()
}
```
