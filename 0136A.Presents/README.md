# 136A. Presents

## Problem

Little Petya very much likes gifts. Recently he has received a new laptop as a New Year gift from his mother. He immediately decided to give it to somebody else as what can be more pleasant than giving somebody gifts. And on this occasion he organized a New Year party at his place and invited n his friends there.

If there's one thing Petya likes more that receiving gifts, that's watching others giving gifts to somebody else. Thus, he safely hid the laptop until the next New Year and made up his mind to watch his friends exchanging gifts while he does not participate in the process. He numbered all his friends with integers from 1 to $n$. Petya remembered that a friend number $i$ gave a gift to a friend number $p_i$. He also remembered that each of his friends received exactly one gift.

Now Petya wants to know for each friend i the number of a friend who has given him a gift.

[Link to CodeForces](https://codeforces.com/problemset/problem/136/A)

## Input

The first line contains one integer $n$ ($1 \leq n \leq 100$) — the quantity of friends Petya invited to the party. The second line contains n space-separated integers: the $i$-th number is $p_i$ — the number of a friend who gave a gift to friend number $i$. It is guaranteed that each friend received exactly one gift. It is possible that some friends do not share Petya's ideas of giving gifts to somebody else. Those friends gave the gifts to themselves.

## Output

Print $n$ space-separated integers: the $i$-th number should equal the number of the friend who gave a gift to friend number $i$.

## Examples

### Input

```
4
2 3 4 1
```

### Output

```
4 1 2 3
```

### Input

```
3
1 3 2
```

### Output

```
1 3 2
```

### Input

```
2
1 2
```

### Output

```
1 2
```

## Solutions

### Constraints

  - Time limit per test: 2 seconds
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

### GO 1.17.5

| Problem |    Lang   |  Verdict | Time  | Memory |
|:-------:|:---------:|:--------:|:-----:|:------:|
| 136A-9  |    Go     | Accepted | 62 ms | 24 KB  |

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
	var friendCount int
	Fscanf("%d\n", &friendCount)

	giftFroms := make([]int, friendCount)

	var giftTo int
	for friendIndex := 0; friendIndex < friendCount; friendIndex++ {
		Fscanf("%d", &giftTo)
		giftFroms[giftTo-1] = friendIndex + 1
	}

	for friendIndex, giftFrom := range giftFroms {
		if friendIndex > 0 {
			Fprintf(" %d", giftFrom)
		} else {
			Fprintf("%d", giftFrom)
		}
	}
	Fprintf("\n")
}

func main() {
	Solution()
	writer.Flush()
}
```
