# 228A. Is your horseshoe on the other hoof?

## Problem

Valera the Horse is going to the party with friends. He has been following the fashion trends for a while, and he knows that it is very popular to wear all horseshoes of different color. Valera has got four horseshoes left from the last year, but maybe some of them have the same color. In this case he needs to go to the store and buy some few more horseshoes, not to lose face in front of his stylish comrades.

Fortunately, the store sells horseshoes of all colors under the sun and Valera has enough money to buy any four of them. However, in order to save the money, he would like to spend as little money as possible, so you need to help Valera and determine what is the minimum number of horseshoes he needs to buy to wear four horseshoes of different colors to a party.

[Link to CodeForces](https://codeforces.com/problemset/problem/228/A)

## Input

The first line contains four space-separated integers $s_1$, $s_2$, $s_3$, $s_4$ ($1 \leq s_1, s_2, s_3, s_4 \leq 109$) — the colors of horseshoes Valera has.

Consider all possible colors indexed with integers.

## Output

Print a single integer — the minimum number of horseshoes Valera needs to buy.

## Examples

### Input

```
1 7 3 3
```

### Output

```
1
```

### Input

```
7 7 7 7
```

### Output

```
3
```

## Solutions

### Constraints

  - Time limit per test: 2 seconds
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

### Go 1.17.5

|  Problem  |    Lang   |  Verdict | Time  | Memory |
|:---------:|:---------:|:--------:|:-----:|:------:|
| 228A - 20 |    Go     | Accepted | 60 ms | 12 KB  |

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
	haveShoes := map[int]bool{}

	var shoe int
	for i := 0; i < 4; i++ {
		Fscanf("%d", &shoe)
		haveShoes[shoe] = true
	}

	Fprintf("%d\n", 4-len(haveShoes))
}

func main() {
	Solution()
	writer.Flush()
}
```
