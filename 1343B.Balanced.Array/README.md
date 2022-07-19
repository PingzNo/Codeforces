# 1343B. Balanced Array

## Problem

You are given a positive integer $n$, it is guaranteed that $n$ is even (i.e. divisible by 2).

You want to construct the array $a$ of length $n$ such that:

  - The first $n/2$ elements of $a$ are even (divisible by 2);
  - The second $n/2$ elements of $a$ are odd (not divisible by 2);
  - All elements of $a$ are distinct and positive;
  - The sum of the first half equals to the sum of the second half ($\sum_{i=1}^{n/2}a_i=\sum_{i=n/2+1}^{n}a_i$).

If there are multiple answers, you can print any. It is not guaranteed that the answer exists.

You have to answer $t$ independent test cases.

[Link to CodeForces](https://codeforces.com/problemset/problem/1343/B)

## Input

The first line of the input contains one integer $t$ ($1 \leq t \leq 10^4$) — the number of test cases. Then $t$ test cases follow.

The only line of the test case contains one integer $n$ ($2 \leq n \leq 2 \cdot 10^5$) — the length of the array. It is guaranteed that that $n$ is even (i.e. divisible by 2).

It is guaranteed that the sum of n over all test cases does not exceed $2 \cdot 10^5$ ($\sum n \leq 2 \cdot 10^5$).

## Output

For each test case, print the answer — "NO" (without quotes), if there is no suitable answer for the given test case or "YES" in the first line and any suitable
array $a_1$, $a_2$, …, $a_n$ ($1 \leq a_i \leq 10^9$) satisfying conditions from the problem statement on the second line.

## Examples

### Input

```
5
2
4
6
8
10
```

### Output

```
NO
YES
2 4 1 5
NO
YES
2 4 6 8 1 3 5 11
NO
```

## Solutions

### Constraints

  - Time limit per test: 10 seconds
  - Memory limit per test: 64 megabytes
  - Input: standard input
  - Output: standard output

### Python 3.8.10

|  Problem   |    Lang   |  Verdict |  Time  |  Memory   |
|:----------:|:---------:|:--------:|:------:|:---------:|
| 1343B - 11 |    Go     | Accepted | 186 ms | 15744 KB  |

[Link to source code](solution.py)

```python
from typing import List

def solution():
    case_count = int(input())

    for _ in range(case_count):
        num_count = int(input())

        half_count = num_count // 2
        if half_count % 2 == 1:
            print("NO")
            continue

        half_half_count = half_count // 2
        ref_num = 0

        result: List[int] = []
        for num_index in range(half_half_count, 0, -1):
            result.append(ref_num - num_index * 2)

        for num_index in range(half_half_count):
            result.append(ref_num + num_index * 2)

        for num_index in range(half_half_count, 0, -1):
            result.append(ref_num - num_index * 2 - 1)

        for num_index in range(half_half_count):
            result.append(ref_num + num_index * 2 + 1)

        first = result[0]
        for num_index in range(len(result)):
            result[num_index] = result[num_index] - first + 2

        print("YES")
        print(" ".join([str(x) for x in result]))


if __name__ == "__main__":
	solution()
```

### Go 1.17.5

|  Problem   |    Lang   |  Verdict |  Time  |  Memory  |
|:----------:|:---------:|:--------:|:------:|:--------:|
| 1343B - 11 |    Go     | Accepted |  62 ms | 5056 KB  |

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
	var caseCount int
	Fscanf("%d\n", &caseCount)

	var numCount int
	for caseIndex := 0; caseIndex < caseCount; caseIndex++ {
		Fscanf("%d\n", &numCount)

		halfCount := numCount / 2
		if halfCount%2 == 1 {
			Fprintf("NO\n")
			continue
		}

		halfHalfCount := halfCount / 2
		refNum := 0

		var result []int
		for numIndex := halfHalfCount; numIndex > 0; numIndex-- {
			result = append(result, refNum-numIndex*2)
		}

		for numIndex := 0; numIndex < halfHalfCount; numIndex++ {
			result = append(result, refNum+numIndex*2)
		}

		for numIndex := halfHalfCount; numIndex > 0; numIndex-- {
			result = append(result, refNum-numIndex*2-1)
		}

		for numIndex := 0; numIndex < halfHalfCount; numIndex++ {
			result = append(result, refNum+numIndex*2+1)
		}

		first := result[0]
		for numIndex := 0; numIndex < len(result); numIndex++ {
			result[numIndex] = result[numIndex] - first + 2
		}

		Fprintf("YES\n")
		for i, x := range result {
			if i > 0 {
				Fprintf(" ")
			}

			Fprintf("%d", x)
		}
		Fprintf("\n")
	}
}

func main() {
	Solution()
	writer.Flush()
}
```
