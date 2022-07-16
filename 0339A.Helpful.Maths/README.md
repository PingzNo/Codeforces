# 339A. Helpful Maths

## Problem

Xenia the beginner mathematician is a third year student at elementary school. She is now learning the addition operation.

The teacher has written down the sum of multiple numbers. Pupils should calculate the sum. To make the calculation easier, the sum only contains numbers 1, 2 and 3. Still, that isn't enough for Xenia. She is only beginning to count, so she can calculate a sum only if the summands follow in non-decreasing order. For example, she can't calculate sum $1+3+2+1$ but she can calculate sums $1+1+2$ and $3+3$.

You've got the sum that was written on the board. Rearrange the summans and print the sum in such a way that Xenia can calculate the sum.

[Link to CodeForces](https://codeforces.com/problemset/problem/339/A)

## Input

The first line contains a non-empty string s — the sum Xenia needs to count. String s contains no spaces. It only contains digits and characters "+". Besides, string s is a correct sum of numbers 1, 2 and 3. String s is at most 100 characters long.

## Output

Print the new sum that Xenia can count.

## Examples

### Input

```
3+2+1
```

### Output

```
1+2+3
```

### Input

```
1+1+3+1+3
```

### Output

```
1+1+1+3+3
```

### Input

```
2
```

### Output

```
2
```

## Constraints

  - Time limit per test: 2 seconds
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

## Solutions

### Go 1.17.5

| Problem |    Lang   |  Verdict | Time  | Memory |
|:-------:|:---------:|:--------:|:-----:|:------:|
| 339A-11 |    Go     | Accepted | 60 ms |  0 KB  |

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
	var line string
	Fscanf("%s\n", &line)

	var counts [3]int
	for _, character := range line {
		if character != '+' {
			counts[character-'1']++
		}
	}

	isFirst := true
	for index := 0; index < 3; index++ {
		for count := 0; count < counts[index]; count++ {
			if isFirst {
				isFirst = false
			} else {
				Fprintf("+")
			}
			Fprintf("%d", index+1)
		}
	}
	Fprintf("\n")
}

func main() {
	Solution()
	writer.Flush()
}
```
