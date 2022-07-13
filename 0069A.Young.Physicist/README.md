# 69A. Young Physicist

## Problem

A guy named Vasya attends the final grade of a high school. One day Vasya decided to watch a match of his favorite hockey team. And, as the boy loves hockey very much, even more than physics, he forgot to do the homework. Specifically, he forgot to complete his physics tasks. Next day the teacher got very angry at Vasya and decided to teach him a lesson. He gave the lazy student a seemingly easy task: You are given an idle body in space and the forces that affect it. The body can be considered as a material point with coordinates (0; 0; 0). Vasya had only to answer whether it is in equilibrium. "Piece of cake" — thought Vasya, we need only to check if the sum of all vectors is equal to 0. So, Vasya began to solve the problem. But later it turned out that there can be lots and lots of these forces, and Vasya can not cope without your help. Help him. Write a program that determines whether a body is idle or is moving by the given vectors of forces.

[Link to CodeForces](https://codeforces.com/problemset/problem/69/A)

## Input

The first line contains a positive integer $n$ ($1 \leq n \leq 100$), then follow $n$ lines containing three integers each: the $x_i$ coordinate, the $y_i$ coordinate and the $z_i$ coordinate of the force vector, applied to the body ( $- 100 \leq x_i, y_i, z_i \leq 100$).

## Output

Print the word "YES" if the body is in equilibrium, or the word "NO" if it is not.

## Examples

### Input

```
3
4 1 7
-2 4 -1
1 -5 -3
```

### Output

```
NO
```

### Input

```
3
3 -1 7
-5 2 -4
2 -1 -3
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

| Problem  |    Lang   |  Verdict | Time  | Memory |
|:--------:|:---------:|:--------:|:-----:|:------:|
| 69A - 23 |    Go     | Accepted | 62 ms | 64 KB  |

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
	var caseCount int
	Fscanf("%d\n", &caseCount)

	var sumX, sumY, sumZ int
	var x, y, z int
	for caseIndex := 0; caseIndex < caseCount; caseIndex++ {
		Fscanf("%d %d %d\n", &x, &y, &z)
		sumX += x
		sumY += y
		sumZ += z
	}

	if sumX == 0 && sumY == 0 && sumZ == 0 {
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
