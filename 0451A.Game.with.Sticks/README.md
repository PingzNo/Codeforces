# 451A. Game With Sticks

## Problem

fter winning gold and silver in IOI 2014, Akshat and Malvika want to have some fun. Now they are playing a game on a grid made of n horizontal and m vertical sticks.

An intersection point is any point on the grid which is formed by the intersection of one horizontal stick and one vertical stick.

In the grid shown below, $n = 3$ and $m = 3$. There are $n + m = 6$ sticks in total (horizontal sticks are shown in red and vertical sticks are shown in green).
There are $n \cdot m = 9$ intersection points, numbered from 1 to 9.

![Illustration](illustration_0.png)

The rules of the game are very simple. The players move in turns. Akshat won gold, so he makes the first move. During his/her move, a player must choose any remaining intersection point and remove from the grid all sticks which pass through this point. A player will lose the game if he/she cannot make a move (i.e. there are no intersection points remaining on the grid at his/her move).

Assume that both players play optimally. Who will win the game?

[Link to CodeForces](https://codeforces.com/problemset/problem/451/A)

## Input

The first line of input contains two space-separated integers, $n$ and $m$ ($1 \leq n$, $m \leq 100$).

## Output

Print a single line containing "Akshat" or "Malvika" (without the quotes), depending on the winner of the game.

## Examples

### Input

```
2 2
```

### Output

```
Malvika
```

### Input

```
2 3
```

### Output

```
Malvika
```

### Input

```
3 3
```

### Output

```
Akshat
```

## Note

Explanation of the first sample:

The grid has four intersection points, numbered from 1 to 4.

![Illustration](illustration_1.png)

If Akshat chooses intersection point 1, then he will remove two sticks (1 - 2 and 1 - 3). The resulting grid will look like this.

![Illustration](illustration_2.png)

Now there is only one remaining intersection point (i.e. 4). Malvika must choose it and remove both remaining sticks. After her move the grid will be empty.

In the empty grid, Akshat cannot make any move, hence he will lose.

Since all 4 intersection points of the grid are equivalent, Akshat will lose no matter which one he picks.

## Solutions

### Constraints

  - Time limit per test: 1 second
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

### Go 1.17.5

|  Problem  |    Lang   |  Verdict | Time  | Memory |
|:---------:|:---------:|:--------:|:-----:|:------:|
| 451A - 29 |    Go     | Accepted | 31 ms | 60 KB  |

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
	var n, m int
	Fscanf("%d %d\n", &n, &m)

	var minValue int
	if n < m {
		minValue = n
	} else {
		minValue = m
	}

	if minValue%2 == 1 {
		Fprintf("Akshat\n")
	} else {
		Fprintf("Malvika\n")
	}
}

func main() {
	Solution()
	writer.Flush()
}
```
