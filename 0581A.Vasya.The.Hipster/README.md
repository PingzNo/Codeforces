# 581A. Vasya the Hipster

## Problem

One day Vasya the Hipster decided to count how many socks he had. It turned out that he had a red socks and b blue socks.

According to the latest fashion, hipsters should wear the socks of different colors: a red one on the left foot, a blue one on the right foot.

Every day Vasya puts on new socks in the morning and throws them away before going to bed as he doesn't want to wash them.

Vasya wonders, what is the maximum number of days when he can dress fashionable and wear different socks, and after that, for how many days he can then wear the same socks until he either runs out of socks or cannot make a single pair from the socks he's got.

Can you help him?

[Link to CodeForces](https://codeforces.com/problemset/problem/581/A)

## Input

The single line of the input contains two positive integers $a$ and $b$ ($1 \leq a, b \leq 100$) — the number of red and blue socks that Vasya's got.

## Output

Print two space-separated integers — the maximum number of days when Vasya can wear different socks and the number of days when he can wear the same socks until he either runs out of socks or cannot make a single pair from the socks he's got.

Keep in mind that at the end of the day Vasya throws away the socks that he's been wearing on that day.

## Examples

### Input

```
3 1
```

### Output

```
1 1
```

### Input

```
2 3
```

### Output

```
2 0
```

### Input

```
7 3
```

### Output

```
3 2
```

## Solutions

### Constraints

  - Time limit per test: 1 second
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

### Go 1.18.3

|   Problem  |    Lang   |  Verdict |  Time  |  Memory  |
|:----------:|:---------:|:--------:|:------:|:--------:|
| 581A - 18  |    Go     | Accepted | 31 ms  |  56  KB  |

[Link to source code](solution.go)

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"math"
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
	var redCount, blueCount float64
	Fscanf("%f %f\n", &redCount, &blueCount)

	fashionCount := int32(math.Min(redCount, blueCount))
	pairCount := (int32(math.Max(redCount, blueCount)) - fashionCount) / 2
	Fprintf("%d %d\n", fashionCount, pairCount)
}

func main() {
	Solution()
	writer.Flush()
}
```
