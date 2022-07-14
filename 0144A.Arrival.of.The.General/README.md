# 144A. Arrival of the General

## Problem

A Ministry for Defense sent a general to inspect the Super Secret Military Squad under the command of the Colonel SuperDuper. Having learned the news, the colonel ordered to all n squad soldiers to line up on the parade ground.

By the military charter the soldiers should stand in the order of non-increasing of their height. But as there's virtually no time to do that, the soldiers lined up in the arbitrary order. However, the general is rather short-sighted and he thinks that the soldiers lined up correctly if the first soldier in the line has the maximum height and the last soldier has the minimum height. Please note that the way other solders are positioned does not matter, including the case when there are several soldiers whose height is maximum or minimum. Only the heights of the first and the last soldier are important.

For example, the general considers the sequence of heights (4, 3, 4, 2, 1, 1) correct and the sequence (4, 3, 1, 2, 2) wrong.

Within one second the colonel can swap any two neighboring soldiers. Help him count the minimum time needed to form a line-up which the general will consider correct.

[Link to CodeForces](https://codeforces.com/problemset/problem/144/A)

## Input

The first input line contains the only integer $n$ ($2 \leq n \leq 100$) which represents the number of soldiers in the line. The second line contains integers $a_1$, $a_2$, ..., $a_n$ ($1 \leq a_i \leq 100$) the values of the soldiers' heights in the order of soldiers' heights' increasing in the order from the beginning of the line to its end. The numbers are space-separated. Numbers $a_1$, $a_2$, ..., $a_n$ are not necessarily different.

## Output

Print the only integer — the minimum number of seconds the colonel will need to form a line-up the general will like.

## Examples

### Input

```
4
33 44 11 22
```

### Output

```
2
```

### Input

```
7
10 10 58 31 63 40 76
```

### Output

```
10
```

## Note

In the first sample the colonel will need to swap the first and second soldier and then the third and fourth soldier. That will take 2 seconds. The resulting position of the soldiers is (44, 33, 22, 11).

In the second sample the colonel may swap the soldiers in the following sequence:

  1. (10, 10, 58, 31, 63, 40, 76)
  2. (10, 58, 10, 31, 63, 40, 76)
  3. (10, 58, 10, 31, 63, 76, 40)
  4. (10, 58, 10, 31, 76, 63, 40)
  5. (10, 58, 31, 10, 76, 63, 40)
  6. (10, 58, 31, 76, 10, 63, 40)
  7. (10, 58, 31, 76, 63, 10, 40)
  8. (10, 58, 76, 31, 63, 10, 40)
  9. (10, 76, 58, 31, 63, 10, 40)
  10. (76, 10, 58, 31, 63, 10, 40)
  11. (76, 10, 58, 31, 63, 40, 10)

## Solutions

### Constraints

  - Time limit per test: 2 seconds
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

### Go 1.17.5

|  Problem  |    Lang   |  Verdict | Time  | Memory |
|:---------:|:---------:|:--------:|:-----:|:------:|
| 144A - 19 |    Go     | Accepted | 62 ms | 64 KB  |

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
	var soldierCount int
	Fscanf("%d\n", &soldierCount)

	minHeight := 101
	minIndex := -1
	maxHeight := 0
	maxIndex := -1

	var height int
	for soldierIndex := 0; soldierIndex < soldierCount; soldierIndex++ {
		Fscanf("%d", &height)
		if height <= minHeight {
			minHeight = height
			minIndex = soldierIndex
		}

		if height > maxHeight {
			maxHeight = height
			maxIndex = soldierIndex
		}
	}

	results := (soldierCount - minIndex) + (maxIndex - 1)
	if minIndex < maxIndex {
		results--
	}

	Fprintf("%d\n", results)
}

func main() {
	Solution()
	writer.Flush()
}
```