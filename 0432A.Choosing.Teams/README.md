# 432A. Choosing Teams

## Problem

The Saratov State University Olympiad Programmers Training Center (SSU OPTC) has n students. For each student you know the number of times he/she has participated in the ACM ICPC world programming championship. According to the ACM ICPC rules, each person can participate in the world championship at most 5 times.

The head of the SSU OPTC is recently gathering teams to participate in the world championship. Each team must consist of exactly three people, at that, any
person cannot be a member of two or more teams. What maximum number of teams can the head make if he wants each team to participate in the world championship
with the same members at least $k$ times?

[Link to CodeForces](https://codeforces.com/problemset/problem/432/A)

## Input

The first line contains two integers, $n$ and $k$ ($1 \leq n \leq 2000$; $1 \leq k \leq 5$). The next line contains $n$ integers: $y_1$, $y_2$, ..., $y_n$ ($0 \leq y_i \leq 5$), where $y_i$ shows the
number of times the $i$-th person participated in the ACM ICPC world championship.

## Output

Print a single number — the answer to the problem.

## Examples

### Input

```
5 2
0 4 5 1 0
```

### Output

```
1
```

### Input

```
6 4
0 1 2 3 4 5
```

### Output

```
0
```

### Input

```
6 5
0 0 0 0 0 0
```

### Output

```
2
```

## Note

In the first sample only one team could be made: the first, the fourth and the fifth participants.

In the second sample no teams could be created.

In the third sample two teams could be created. Any partition into two teams fits.

## Solutions

### Constraints

  - Time limit per test: 1 second
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

### Go 1.18.3

|  Problem  |    Lang   |  Verdict |  Time  |  Memory  |
|:---------:|:---------:|:--------:|:------:|:--------:|
| 432A - 17 |    Go     | Accepted |  31 ms |  64  KB  |

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
	var personCount, minContestCount int
	Fscanf("%d %d\n", &personCount, &minContestCount)

	var contestCount, candidateCount int
	for personIndex := 0; personIndex < personCount; personIndex++ {
		Fscanf("%d", &contestCount)
		if 5-contestCount >= minContestCount {
			candidateCount++
		}
	}
	Fscanf("\n")

	Fprintf("%d\n", candidateCount/3)
}

func main() {
	Solution()
	writer.Flush()
}
```
