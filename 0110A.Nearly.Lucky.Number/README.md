# 110A. Nearly Lucky Number

## Problem

Petya loves lucky numbers. We all know that lucky numbers are the positive integers whose decimal representations contain only the lucky digits 4 and 7. For example, numbers 47, 744, 4 are lucky and 5, 17, 467 are not.

Unfortunately, not all numbers are lucky. Petya calls a number nearly lucky if the number of lucky digits in it is a lucky number. He wonders whether number n is a nearly lucky number.

[Link to CodeForces](https://codeforces.com/problemset/problem/110/A)

## Input

The only line contains an integer $n$ ($1 \leq n \leq 10^{18}$).

Please do not use the %lld specificator to read or write 64-bit numbers in С++. It is preferred to use the cin, cout streams or the %I64d specificator.

## Output

Print on the single line "YES" if n is a nearly lucky number. Otherwise, print "NO" (without the quotes).


## Examples

### Input

```
40047
```

### Output

```
NO
```

### Input

```
7747774
```

### Output

```
YES
```

### Input

```
1000000000000000000
```

### Output

```
NO
```

## Note

In the first sample there are 3 lucky digits (first one and last two), so the answer is "NO".

In the second sample there are 7 lucky digits, 7 is lucky number, so the answer is "YES".

In the third sample there are no lucky digits, so the answer is "NO".

## Solutions

### Constraints

  - Time limit per test: 2 seconds
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

### Go 1.17.5

|  Problem  |    Lang   |  Verdict | Time  | Memory |
|:---------:|:---------:|:--------:|:-----:|:------:|
| 110A - 20 |     Go    | Accepted | 62 ms |  16 KB |

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
	var number string
	Fscanf("%s\n", &number)

	luckyCount := 0
	for _, r := range number {
		if r == '4' || r == '7' {
			luckyCount++
		}
	}

	isNearlyLucky := luckyCount > 0

	for luckyCount > 0 {
		remainder := luckyCount % 10
		if remainder != 4 && remainder != 7 {
			isNearlyLucky = false
			break
		}

		luckyCount /= 10
	}

	if isNearlyLucky {
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
