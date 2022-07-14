# 148A. Insomnia Cure

## Problem

«One dragon. Two dragon. Three dragon», — the princess was counting. She had trouble falling asleep, and she got bored of counting lambs when she was nine.

However, just counting dragons was boring as well, so she entertained herself at best she could. Tonight she imagined that all dragons were here to steal her,
and she was fighting them off. Every $k$-th dragon got punched in the face with a frying pan. Every $l$-th dragon got his tail shut into the balcony door. Every
$m$-th dragon got his paws trampled with sharp heels. Finally, she threatened every $n$-th dragon to call her mom, and he withdrew in panic.

How many imaginary dragons suffered moral or physical damage tonight, if the princess counted a total of $d$ dragons?

[Link to CodeForces](https://codeforces.com/problemset/problem/148/A)

## Input

Input data contains integer numbers $k$, $l$, $m$, $n$ and $d$, each number in a separate line ($1 \leq k, l, m, n \leq 10$, $1 \leq d \leq 105$).

## Output

Output the number of damaged dragons.

## Note

In the first case every first dragon got punched with a frying pan. Some of the dragons suffered from other reasons as well, but the pan alone would be enough.

In the second case dragons 1, 7, 11, 13, 17, 19 and 23 escaped unharmed.

## Examples

### Input

```
1
2
3
4
12
```

### Output

```
12
```

### Input

```
2
3
4
5
24
```

### Output

```
17
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
| 148A - 10 |    Go     | Accepted | 60 ms | 68 KB  |

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
	var divider, dragonCount int
	dividers := []int{}
	for i := 0; i < 4; i++ {
		Fscanf("%d\n", &divider)
		if divider == 1 {
			for j := i + 1; j < 5; j++ {
				Fscanf("%d\n", &dragonCount)
			}

			Fprintf("%d\n", dragonCount)
			return
		}

		isDivided := false
		for j, x := range dividers {
			if x%divider == 0 {
				dividers[j] = divider
				isDivided = true
				continue
			}

			if divider%x == 0 {
				isDivided = true
				continue
			}
		}

		if !isDivided {
			dividers = append(dividers, divider)
		}
	}

	Fscanf("%d\n", &dragonCount)

	damagedCount := 0
	for x := 1; x <= dragonCount; x++ {
		for _, y := range dividers {
			if x%y == 0 {
				damagedCount++
				break
			}
		}
	}

	Fprintf("%d\n", damagedCount)
}

func main() {
	Solution()
	writer.Flush()
}
```
