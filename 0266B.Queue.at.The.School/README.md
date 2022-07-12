# 266B. Queue at the School

## Problem

During the break the schoolchildren, boys and girls, formed a queue of n people in the canteen. Initially the children stood in the order they entered the canteen. However, after a while the boys started feeling awkward for standing in front of the girls in the queue and they started letting the girls move forward each second.

Let's describe the process more precisely. Let's say that the positions in the queue are sequentially numbered by integers from 1 to $n$, at that the person in the position number 1 is served first. Then, if at time x a boy stands on the $i$-th position and a girl stands on the $(i + 1)$-th position, then at time $x + 1$ the $i$-th position will have a girl and the $(i + 1)$-th position will have a boy. The time is given in seconds.

You've got the initial position of the children, at the initial moment of time. Determine the way the queue is going to look after t seconds.

[Link to CodeForces](https://codeforces.com/problemset/problem/266/B)

## Input

The first line contains two integers $n$ and $t$ ($1 \leq n, t \leq 50$), which represent the number of children in the queue and the time after which the queue will transform into the arrangement you need to find.

The next line contains string $s$, which represents the schoolchildren's initial arrangement. If the $i$-th position in the queue contains a boy, then the $i$-th
character of string $s$ equals "B", otherwise the $i$-th character equals "G".

## Output

Print string $a$, which describes the arrangement after $t$ seconds. If the $i$-th position has a boy after the needed time, then the $i$-th character a must equal "B", otherwise it must equal "G".

## Examples

### Input

```
5 1
BGGBG
```

### Output

```
GBGGB
```

### Input

```
5 2
BGGBG
```

### Output

```
GGBGB
```

### Input

```
4 1
GGGB
```

### Output

```
GGGB
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
| 266B - 10 |    Go     | Accepted | 62 ms |  60 KB  |

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
	var headCount, shiftCount int
	Fscanf("%d %d\n", &headCount, &shiftCount)

	var line string
	Fscanf("%s\n", &line)

	firstQueue := []rune(line)
	secondQueue := []rune(line)

	p := &firstQueue
	q := &secondQueue

	for shiftIndex := 0; shiftIndex < shiftCount; shiftIndex++ {
		queueIndex := 1
		for queueIndex < headCount {
			if (*p)[queueIndex-1] == 'B' && (*p)[queueIndex] == 'G' {
				(*q)[queueIndex-1], (*q)[queueIndex] = 'G', 'B'
				if queueIndex < headCount-1 && queueIndex+2 >= headCount {
					(*q)[headCount-1] = (*p)[headCount-1]
				}
				queueIndex += 2
			} else {
				(*q)[queueIndex-1], (*q)[queueIndex] = (*p)[queueIndex-1], (*p)[queueIndex]
				queueIndex++
			}
		}

		p, q = q, p
	}

	for _, r := range *p {
		Fprintf("%c", r)
	}
	Fprintf("\n")
}

func main() {
	Solution()
	writer.Flush()
}
```
