# 200B. Drinks

## Problem

Little Vasya loves orange juice very much. That's why any food and drink in his kitchen necessarily contains orange juice. There are n drinks in his fridge, the volume fraction of orange juice in the $i$-th drink equals pi percent.

One day Vasya decided to make himself an orange cocktail. He took equal proportions of each of the n drinks and mixed them. Then he wondered, how much orange juice the cocktail has.

Find the volume fraction of orange juice in the final drink.

[Link to CodeForces](https://codeforces.com/problemset/problem/200/B)

## Input

The first input line contains a single integer $n$ ($1 \leq n \leq 100$) — the number of orange-containing drinks in Vasya's fridge. The second line contains $n$ integers $p_i$ ($0 \leq p_i \leq 100$) — the volume fraction of orange juice in the $i$-th drink, in percent. The numbers are separated by a space.

## Output

Print the volume fraction in percent of orange juice in Vasya's cocktail. The answer will be considered correct if the absolute or relative error does not
exceed $10^{-4}$.

## Note

Note to the first sample: let's assume that Vasya takes $x$ milliliters of each drink from the fridge. Then the volume of pure juice in the cocktail will equal
$\dfrac{x}{2} + \dfrac{x}{2} + x = 2 \cdot x$ milliliters. The total cocktail's volume equals $3 \cdot x$ milliliters, so the volume fraction of the juice in the cocktail
equals $\dfrac{2 \cdot x}{3 \cdot x} = \dfrac{2}{3}$, that is, 66.(6) percent.

## Examples

### Input

```
3
50 50 100
```

### Output

```
66.666666666667
```

### Input

```
4
0 25 50 75
```

### Output

```
37.500000000000
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
| 200B - 10 |    Go     | Accepted | 60 ms | 64 KB  |

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
	var bottleCount int
	Fscanf("%d\n", &bottleCount)

	var percent, percentSum int
	for bottleIndex := 0; bottleIndex < bottleCount; bottleIndex++ {
		Fscanf("%d", &percent)
		percentSum += percent
	}

	Fprintf("%.5f\n", float64(percentSum)/float64(bottleCount))
}

func main() {
	Solution()
	writer.Flush()
}
```
