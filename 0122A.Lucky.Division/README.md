# 122A. Lucky Division

## Problem

Petya loves lucky numbers. Everybody knows that lucky numbers are positive integers whose decimal representation contains only the lucky digits 4 and 7. For example, numbers 47, 744, 4 are lucky and 5, 17, 467 are not.

Petya calls a number almost lucky if it could be evenly divided by some lucky number. Help him find out if the given number n is almost lucky.

[Link to CodeForces](https://codeforces.com/problemset/problem/122/A)

## Input

The single line contains an integer $n$ ($1 \leq n \leq 1000$) — the number that needs to be checked.

## Output

In the only line print "YES" (without the quotes), if number n is almost lucky. Otherwise, print "NO" (without the quotes).

## Examples

### Input

```
47
```

### Output

```
YES
```

### Input

```
16
```

### Output

```
YES
```

### Input

```
78
```

### Output

```
NO
```

## Note

Note that all lucky numbers are almost lucky as any number is evenly divisible by itself.

In the first sample 47 is a lucky number. In the second sample 16 is divisible by 4.

## Solutions

### Constraints

  - Time limit per test: 2 seconds
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

### Go 1.17.5

| Problem  |    Lang   |  Verdict | Time  | Memory |
|:--------:|:---------:|:--------:|:-----:|:------:|
| 122A - 9 |    Go     | Accepted | 60 ms |  0 KB  |

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
	var number int
	Fscanf("%d\n", &number)

	dividers := []int{4, 7, 47, 74, 447, 474, 477, 744, 747, 774}
	for _, divider := range dividers {
		if divider <= number {
			if number%divider == 0 {
				Fprintf("YES\n")
				return
			}
		} else {
			break
		}
	}

	Fprintf("NO\n")
}

func main() {
	Solution()
	writer.Flush()
}
```
