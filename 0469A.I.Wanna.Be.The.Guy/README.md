# 469A. I Wanna Be the Guy

## Problem

There is a game called "I Wanna Be the Guy", consisting of $n$ levels. Little X and his friend Little Y are addicted to the game. Each of them wants to pass the whole game.

Little X can pass only $p$ levels of the game. And Little Y can pass only $q$ levels of the game. You are given the indices of levels Little X can pass and the indices of levels Little Y can pass. Will Little X and Little Y pass the whole game, if they cooperate each other?

[Link to CodeForces](https://codeforces.com/problemset/problem/469/A)

## Input

The first line contains a single integer $n$ ($1 \leq n \leq 100$).

The next line contains an integer $p$ ($0 \leq p \leq n$) at first, then follows $p$ distinct integers $a_1$, $a_2$, ..., $a_p$ ($1 \leq a_i \leq n$). These integers denote the indices of levels Little X can pass. The next line contains the levels Little Y can pass in the same format. It's assumed that levels are numbered from 1 to $n$.

## Output

If they can pass all the levels, print "I become the guy.". If it's impossible, print "Oh, my keyboard!" (without the quotes).

## Examples

### Input

```
4
3 1 2 3
2 2 4
```

### Output

```
I become the guy.
```

### Input

```
4
3 1 2 3
2 2 3
```

### Output

```
Oh, my keyboard!
```

## Solutions

### Constraints

  - Time limit per test: 1 second
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

### GNU C++17 7.3.0

|  Problem  |    Lang   |  Verdict | Time  | Memory |
|:---------:|:---------:|:--------:|:-----:|:------:|
| 469A - 24 |    Go     | Accepted | 31 ms | 84 KB  |

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
	var levelCount int
	Fscanf("%d\n", &levelCount)

	levelPassed := map[int]bool{}

	var playCount, level int
	for personIndex := 0; personIndex < 2; personIndex++ {
		Fscanf("%d", &playCount)
		for playIndex := 0; playIndex < playCount; playIndex++ {
			Fscanf("%d", &level)
			levelPassed[level-1] = true
		}

		Fscanf("\n")
	}

	if len(levelPassed) == levelCount {
		Fprintf("I become the guy.\n")
	} else {
		Fprintf("Oh, my keyboard!\n")
	}

}

func main() {
	Solution()
	writer.Flush()
}
```
