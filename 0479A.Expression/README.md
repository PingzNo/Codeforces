# 479A. Expression

## Problem

Petya studies in a school and he adores Maths. His class has been studying arithmetic expressions. On the last class the teacher wrote three positive integers $a$, $b$, $c$ on the blackboard. The task was to insert signs of operations '+' and '*', and probably brackets between the numbers so that the value of the resulting expression is as large as possible. Let's consider an example: assume that the teacher wrote numbers 1, 2 and 3 on the blackboard. Here are some ways of placing signs and brackets:

  - 1+2*3=7
  - 1*(2+3)=5
  - 1*2*3=6
  - (1+2)*3=9

Note that you can insert operation signs only between $a$ and $b$, and between $b$ and $c$, that is, you cannot swap integers. For instance, in the given sample you cannot get expression (1+3)*2.

It's easy to see that the maximum value that you can obtain is 9.

Your task is: given $a$, $b$ and $c$ print the maximum value that you can get.

[Link to CodeForces](https://codeforces.com/problemset/problem/479/A)

## Input

The input contains three integers $a$, $b$ and $c$, each on a single line ($1 \leq a, b, c \leq 10$).

## Output

Print the maximum value of the expression that you can obtain.

## Examples

### Input

```
1
2
3
```

### Output

```
9
```

### Input

```
2
10
3
```

### Output

```
60
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
| 479A - 26 |    Go     | Accepted | 15 ms | 64 KB  |

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
	var a, b, c int
	Fscanf("%d\n%d\n%d\n", &a, &b, &c)

	var maxValue, result int

	result = a + b + c
	if result > maxValue {
		maxValue = result
	}

	result = a * b * c
	if result > maxValue {
		maxValue = result
	}

	result = (a + b) * c
	if result > maxValue {
		maxValue = result
	}

	result = a * (b + c)
	if result > maxValue {
		maxValue = result
	}

	result = a*b + c
	if result > maxValue {
		maxValue = result
	}

	result = a + b*c
	if result > maxValue {
		maxValue = result
	}

	Fprintf("%d\n", maxValue)
}

func main() {
	Solution()
	writer.Flush()
}
```
