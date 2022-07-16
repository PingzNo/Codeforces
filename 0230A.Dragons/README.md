# 230A. Dragons

## Problem

Kirito is stuck on a level of the MMORPG he is playing now. To move on in the game, he's got to defeat all n dragons that live on this level. Kirito and the dragons have strength, which is represented by an integer. In the duel between two opponents the duel's outcome is determined by their strength. Initially, Kirito's strength equals s.

If Kirito starts duelling with the $i$-th ($1 \leq i \leq n$) dragon and Kirito's strength is not greater than the dragon's strength $x_i$, then Kirito loses the duel and dies. But if Kirito's strength is greater than the dragon's strength, then he defeats the dragon and gets a bonus strength increase by $y_i$.

Kirito can fight the dragons in any order. Determine whether he can move on to the next level of the game, that is, defeat all dragons without a single loss.

[Link to CodeForces](https://codeforces.com/problemset/problem/230/A)

## Input

The first line contains two space-separated integers $s$ and $n$ ($1 \leq s \leq 10^4$, $1 \leq n \leq 10^3$). Then $n$ lines follow: the $i$-th line contains space-separated integers $x_i$ and $y_i$ ($1 \leq x_i \leq 10^4$, $0 \leq y_i \leq 10^4$) — the $i$-th dragon's strength and the bonus for defeating it.

## Output

On a single line print "YES" (without the quotes), if Kirito can move on to the next level and print "NO" (without the quotes), if he can't.

## Examples

### Input

```
2 2
1 99
100 0
```

### Output

```
YES
```

### Input

```
10 1
100 100
```

### Output

```
NO
```

## Note

In the first sample Kirito's strength initially equals 2. As the first dragon's strength is less than 2, Kirito can fight it and defeat it. After that he gets the bonus and his strength increases to 2 + 99 = 101. Now he can defeat the second dragon and move on to the next level.

In the second sample Kirito's strength is too small to defeat the only dragon and win.

## Solutions

### Constraints

  - Time limit per test: 2 seconds
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

### Go 1.17.5

|  Problem  |    Lang   |  Verdict | Time  | Memory |
|:---------:|:---------:|:--------:|:-----:|:------:|
| 230A - 52 |    Go     | Accepted | 62 ms | 60 KB  |

[Link to source code](solution.go)

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Fprintf(format string, x ...interface{}) {
	fmt.Fprintf(writer, format, x...)
}

func Fscanf(format string, x ...interface{}) {
	fmt.Fscanf(reader, format, x...)
}

type Dragon struct {
	Strength int
	Bonus    int
}

func Solution() {
	var strength, dragonCount int
	Fscanf("%d %d\n", &strength, &dragonCount)

	dragons := make([]Dragon, dragonCount)
	for dragonIndex := 0; dragonIndex < dragonCount; dragonIndex++ {
		Fscanf("%d %d\n", &dragons[dragonIndex].Strength, &dragons[dragonIndex].Bonus)
	}

	sort.Slice(dragons, func(i, j int) bool {
		return dragons[i].Strength < dragons[j].Strength
	})

	gamePassed := true
	for dragonIndex := 0; dragonIndex < dragonCount; dragonIndex++ {
		if strength <= dragons[dragonIndex].Strength {
			gamePassed = false
			break
		}

		strength += dragons[dragonIndex].Bonus
	}

	if gamePassed {
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
