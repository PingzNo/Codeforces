# 467A. George and Accommodation

## Problem

George has recently entered the BSUCP (Berland State University for Cool Programmers). George has a friend Alex who has also entered the university. Now they are moving into a dormitory.

George and Alex want to live in the same room. The dormitory has n rooms in total. At the moment the $i$-th room has pi people living in it and the room can
accommodate $q_i$ people in total ($p_i \leq q_i$). Your task is to count how many rooms has free place for both George and Alex.

[Link to CodeForces](https://codeforces.com/problemset/problem/467/A)

## Input

The first line contains a single integer $n$ ($1 \leq n \leq 100$) — the number of rooms.

The $i$-th of the next $n$ lines contains two integers $p_i$ and $q_i$ ($0 \leq p_i \leq q_i \leq 100$) — the number of people who already live in the $i$-th room and the room's capacity.

## Output

Print a single integer — the number of rooms where George and Alex can move in.

## Examples

### Input

```
3
1 1
2 2
3 3
```

### Output

```
0
```

### Input

```
3
1 10
0 10
10 10
```

### Output

```
2
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
| 467A - 17 |    Go     | Accepted | 31 ms | 12 KB  |

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
	var rooCount int
	Fscanf("%d\n", &rooCount)

	var roomToMoveIn int
	var peopleCount, roomCap int
	for roomIndex := 0; roomIndex < rooCount; roomIndex++ {
		Fscanf("%d %d\n", &peopleCount, &roomCap)
		if roomCap-peopleCount >= 2 {
			roomToMoveIn++
		}
	}

	Fprintf("%d\n", roomToMoveIn)
}

func main() {
	Solution()
	writer.Flush()
}
```
