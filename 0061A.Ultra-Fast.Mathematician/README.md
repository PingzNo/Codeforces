# 61A. Ultra-Fast Mathematician

## Problem

Shapur was an extremely gifted student. He was great at everything including Combinatorics, Algebra, Number Theory, Geometry, Calculus, etc. He was not only smart but extraordinarily fast! He could manage to sum $10^{18}$ numbers in a single second.

One day in 230 AD Shapur was trying to find out if any one can possibly do calculations faster than him. As a result he made a very great contest and asked every one to come and take part.

In his contest he gave the contestants many different pairs of numbers. Each number is made from digits 0 or 1. The contestants should write a new number corresponding to the given pair of numbers. The rule is simple: The i-th digit of the answer is 1 if and only if the $i$-th digit of the two given numbers differ. In the other case the i-th digit of the answer is 0.

Shapur made many numbers and first tried his own speed. He saw that he can perform these operations on numbers of length $\infty$ (length of a number is number of digits in it) in a glance! He always gives correct answers so he expects the contestants to give correct answers, too. He is a good fellow so he won't give anyone very big numbers and he always gives one person numbers of same length.

Now you are going to take part in Shapur's contest. See if you are faster and more accurate.

[Link to CodeForces](https://codeforces.com/problemset/problem/61/A)

## Input

There are two lines in each input. Each of them contains a single number. It is guaranteed that the numbers are made from 0 and 1 only and that their length is same. The numbers may start with 0. The length of each number doesn't exceed 100.

## Output

Write one line — the corresponding answer. Do not omit the leading 0s.

## Examples

### Input

```
1010100
0100101
```

### Output

```
1110001
```

### Input

```
000
111
```

### Output

```
111
```

### Input

```
1110
1010
```

### Output

```
0100
```

### Input

```
01110
01100
```

### Output

```
00010
```

## Constraints

  - Time limit per test: 2 seconds
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

## Solutions

### Go 1.17.5

| Problem  |    Lang   |  Verdict | Time  | Memory |
|:--------:|:---------:|:--------:|:-----:|:------:|
| 61A - 20 |    Go     | Accepted | 31 ms | 64 KB  |

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
	var x, y string
	Fscanf("%s\n%s\n", &x, &y)

	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			Fprintf("1")
		} else {
			Fprintf("0")
		}
	}
	Fprintf("\n")
}

func main() {
	Solution()
	writer.Flush()
}
```
