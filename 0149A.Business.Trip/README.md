# 149A. Business trip

## Problem

What joy! Petya's parents went on a business trip for the whole year and the playful kid is left all by himself. Petya got absolutely happy. He jumped on the bed and threw pillows all day long, until...

Today Petya opened the cupboard and found a scary note there. His parents had left him with duties: he should water their favourite flower all year, each day, in the morning, in the afternoon and in the evening. "Wait a second!" — thought Petya. He know for a fact that if he fulfills the parents' task in the $i$-th ($1 \leq i \leq 12$) month of the year, then the flower will grow by ai centimeters, and if he doesn't water the flower in the $i$-th month, then the flower won't grow this month. Petya also knows that try as he might, his parents won't believe that he has been watering the flower if it grows strictly less than by $k$ centimeters.

Help Petya choose the minimum number of months when he will water the flower, given that the flower should grow no less than by $k$ centimeters.

[Link to CodeForces](https://codeforces.com/problemset/problem/149/A)

## Input

The first line contains exactly one integer $k$ ($0 \leq k \leq 100$). The next line contains twelve space-separated integers: the $i$-th ($1 \leq i \leq 12$) number in the line represents $a_i$ ($0 \leq a_i \leq 100$).

## Output

Print the only integer — the minimum number of months when Petya has to water the flower so that the flower grows no less than by $k$ centimeters. If the flower can't grow by $k$ centimeters in a year, print -1.

## Examples

### Input

```
5
1 1 1 1 2 2 3 2 2 1 1 1
```

### Output

```
2
```

### Input

```
0
0 0 0 0 0 0 0 1 1 2 3 0
```

### Output

```
0
```

### Input

```
11
1 1 4 1 1 5 1 1 4 1 1 1
```

### Output

```
3
```

## Note

Let's consider the first sample test. There it is enough to water the flower during the seventh and the ninth month. Then the flower grows by exactly five centimeters.

In the second sample Petya's parents will believe him even if the flower doesn't grow at all ($k = 0$). So, it is possible for Petya not to water the flower at all.

## Solutions

### Constraints

  - Time limit per test: 2 seconds
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

### Go 1.18.3

|  Problem  |    Lang   |  Verdict |  Time  |  Memory  |
|:---------:|:---------:|:--------:|:------:|:--------:|
| 149A - 28 |    Go     | Accepted |  62 ms |  64  KB  |

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

func Solution() {
	var centimeters int
	Fscanf("%d\n", &centimeters)

	nums := make([]int, 12)
	for numIndex := 0; numIndex < len(nums); numIndex++ {
		Fscanf("%d", &nums[numIndex])
	}

	sort.Ints(nums)

	var monthCount int
	for numIndex := len(nums) - 1; numIndex >= 0; numIndex-- {
		if centimeters <= 0 {
			break
		}

		centimeters -= nums[numIndex]
		monthCount++
	}

	if centimeters > 0 {
		Fprintf("-1\n")
	} else {
		Fprintf("%d\n", monthCount)
	}
}

func main() {
	Solution()
	writer.Flush()
}
```
