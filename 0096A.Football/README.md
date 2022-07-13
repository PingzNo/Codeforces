# 96A. Football

## Problem

Petya loves football very much. One day, as he was watching a football match, he was writing the players' current positions on a piece of paper. To simplify the
situation he depicted it as a string consisting of zeroes and ones. A zero corresponds to players of one team; a one corresponds to players of another team. If
there are at least 7 players of some team standing one after another, then the situation is considered dangerous. For example, the situation 00100110111111101
is dangerous and 11110111011101 is not. You are given the current situation. Determine whether it is dangerous or not.

[Link to CodeForces](https://codeforces.com/problemset/problem/96/A)

## Input

The first input line contains a non-empty string consisting of characters "0" and "1", which represents players. The length of the string does not exceed 100 characters. There's at least one player from each team present on the field.

## Output

Print "YES" if the situation is dangerous. Otherwise, print "NO".

## Examples

### Input

```
001001
```

### Output

```
NO
```

### Input

```
1000000001
```

### Output

```
YES
```

## Solutions

### Constraints

  - Time limit per test: 2 seconds
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

### Go 1.17.5

|  Problem |    Lang   |  Verdict | Time  | Memory |
|:--------:|:---------:|:--------:|:-----:|:------:|
| 96A - 10 |     Go    | Accepted | 62 ms | 64 KB  |

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
	var status string
	Fscanf("%s\n", &status)

	team := '0'
	count := 0
	for _, r := range status {
		if r == team {
			count++
		} else {
			team = r
			count = 1
		}

		if count >= 7 {
			break
		}
	}

	if count >= 7 {
		Fprintf("YES\n")
	} else {
		Fprintf("NO\n")
	}
}

func main() {
	Solution()
	writer.Flush()
}

```
