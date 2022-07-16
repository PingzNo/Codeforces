# 158B. counter[0] -= 2

## Problem

After the lessons $n$ groups of schoolchildren went outside and decided to visit Polycarpus to celebrate his birthday. We know that the $i$-th group consists of $s_i$ friends ($1 \leq s_i \leq 4$), and they want to go to Polycarpus together. They decided to get there by taxi. Each car can carry at most four passengers. What minimum number of cars will the children need if all members of each group should ride in the same taxi (but one taxi can take more than one group)?

[Link to CodeForces](https://codeforces.com/problemset/problem/158/B)

## Input

The first line contains integer $n$ ($1 \leq n \leq 10^5$) — the number of groups of schoolchildren. The second line contains a sequence of integers $s_1$, $s_2$, ..., $s_n$ ($1 \leq s_i \leq 4$). The integers are separated by a space, $s_i$ is the number of children in the $i$-th group.

## Output

Print the single number — the minimum number of taxis necessary to drive all children to Polycarpus.

## Examples

### Input

```
5
1 2 4 3 3
```

### Output

```
4
```

### Input

```
8
2 3 4 4 2 1 3 1
```

### Output

```
5
```

## Note

In the first test we can sort the children into four cars like this:

  - The third group (consisting of four children),
  - The fourth group (consisting of three children),
  - The fifth group (consisting of three children),
  - The first and the second group (consisting of one and two children, correspondingly).

There are other ways to sort the groups into four cars.

## Solutions

### Constraints

  - Time limit per test: 3 seconds
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

### Go 1.17.5

|  Problem  |    Lang   |  Verdict | Time  | Memory |
|:---------:|:---------:|:--------:|:-----:|:------:|
| 158B - 10 |    Go     | Accepted | 92 ms | 60 KB  |

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
	var group, groupCount int
	Fscanf("%d\n", &groupCount)

	counter := [4]int{}
	for groupIndex := 0; groupIndex < groupCount; groupIndex++ {
		Fscanf("%d", &group)
		counter[group-1]++
	}

	taxiCount := counter[3]

	taxiCount += counter[2]
	if counter[2] <= counter[0] {
		counter[0] -= counter[2]
	} else {
		counter[0] = 0
	}

	taxiCount += counter[1] / 2
	counter[1] %= 2
	if counter[1] == 1 {
		taxiCount++
		if counter[0] >= 2 {
			counter[0] -= 2
		} else if counter[0] >= 1 {
			counter[0] -= 1
		}
	}

	taxiCount += counter[0] / 4
	if counter[0]%4 > 0 {
		taxiCount++
	}

	Fprintf("%d\n", taxiCount)
}

func main() {
	Solution()
	writer.Flush()
}
```
