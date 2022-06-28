# 791A. Bear and Big Brother

## Problem

Bear Limak wants to become the largest of bears, or at least to become larger than his brother Bob.

Right now, Limak and Bob weigh a and b respectively. It's guaranteed that Limak's weight is smaller than or equal to his brother's weight.

Limak eats a lot and his weight is tripled after every year, while Bob's weight is doubled after every year.

After how many full years will Limak become strictly larger (strictly heavier) than Bob?

[Link to CodeForces](https://codeforces.com/problemset/problem/791/A)

## Input

The only line of the input contains two integers $a$ and $b$ ($1 \leq a \leq b \leq 10$) — the weight of Limak and the weight of Bob respectively.

## Output

Print one integer, denoting the integer number of years after which Limak will become strictly larger than Bob.

## Examples

### Input

```
4 7
```

### Output

```
2
```

### Input

```
4 9
```

### Output

```
3
```

### Input

```
1 1
```

### Output

```
1
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
| 791A - 11 |    Go     | Accepted | 31 ms | 68 KB  |

[Link to source code](solution.go)

```go
package main

import (
	"bufio"
	"fmt"
	"math"
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
	var limak, bob float64
	Fscanf("%f %f\n", &limak, &bob)

	result := (math.Log(bob) - math.Log(limak)) / (math.Log(3) - math.Log(2))
	if result == float64(int(result)) {
		result++
	} else {
		result = float64(int(result) + 1)
	}

	Fprintf("%d\n", int(result))
}

func main() {
	Solution()
	writer.Flush()
}
```
