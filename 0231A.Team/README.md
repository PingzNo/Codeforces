# 231A. Team

## Problem

One day three best friends Petya, Vasya and Tonya decided to form a team and take part in programming contests. Participants are usually offered several problems during programming contests. Long before the start the friends decided that they will implement a problem if at least two of them are sure about the solution. Otherwise, the friends won't write the problem's solution.

This contest offers $n$ problems to the participants. For each problem we know, which friend is sure about the solution. Help the friends find the number of problems for which they will write a solution.

[Link to CodeForces](https://codeforces.com/problemset/problem/231/A)

## Input

The first input line contains a single integer $n$ ($1 \leq n \leq 1000$) — the number of problems in the contest. Then n lines contain three integers each, each integer
is either 0 or 1. If the first number in the line equals 1, then Petya is sure about the problem's solution, otherwise he isn't sure. The second number shows
Vasya's view on the solution, the third number shows Tonya's view. The numbers on the lines are separated by spaces.

## Output

Print a single integer — the number of problems the friends will implement on the contest.

## Examples

### Input

```
3
1 1 0
1 1 1
1 0 0
```

### Output

```
2
```

### Input

```
2
1 0 0
0 1 1
```

### Output

```
1
```

## Note

In the first sample Petya and Vasya are sure that they know how to solve the first problem and all three of them know how to solve the second problem. That means that they will write solutions for these problems. Only Petya is sure about the solution for the third problem, but that isn't enough, so the friends won't take it.

In the second sample the friends will only implement the second problem, as Vasya and Tonya are sure about the solution.

## Constraints

  - Time limit per test: 2 seconds
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

## Solutions

### GNU C++17 7.3.0

| Problem |    Lang   |  Verdict | Time  | Memory |
|:-------:|:---------:|:--------:|:-----:|:------:|
| 231A-11 | GNU C++17 | Accepted | 30 ms |  8 KB  |

[Link to source code](solution.cpp)

```c++
#include <iostream>
#include <string>


void solution() {
    int problem_count = 0;
    std::cin >> problem_count;
    std::cin.ignore();

    std::string line;

    int solved_count = 0;
    for (int problem_index = 0; problem_index < problem_count; ++problem_index) {
        std::getline(std::cin, line);
        solved_count += (line[0] == '1') + (line[2] == '1') + (line[4] == '1') >= 2 ? 1 : 0;
    }
    std::cout << solved_count << '\n';
}


void setup() {
    std::ios_base::sync_with_stdio(false);
    std::cin.tie(NULL);
}


int main() {
    setup();
    solution();
}
```

### Python 3.8.10

| Problem |    Lang   |  Verdict | Time  | Memory |
|:-------:|:---------:|:--------:|:-----:|:------:|
| 231A-11 |  Python 3 | Accepted | 92 ms |  0 KB  |

[Link to source code](solution.py)

```python
def solution():
	problem_count = int(input())

	solved_count = 0
	for _ in range(problem_count):
		petya, vasya, tonya = [int(x) for x in input().split()]
		if petya + vasya + tonya >= 2:
			solved_count += 1

	print(solved_count)


if __name__ == "__main__":
	solution()
```

### Go 1.17.5

| Problem |    Lang   |  Verdict | Time  | Memory |
|:-------:|:---------:|:--------:|:-----:|:------:|
| 231A-11 |    Go     | Accepted | 60 ms |  0 KB  |

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
	var problemCount int
	Fscanf("%d\n", &problemCount)

	solvedCount := 0
	var petya, vasya, tonya int
	for problemIndex := 0; problemIndex < problemCount; problemIndex++ {
		Fscanf("%d %d %d\n", &petya, &vasya, &tonya)
		if petya+vasya+tonya >= 2 {
			solvedCount++
		}
	}

	Fprintf("%d\n", solvedCount)
}

func main() {
	Solution()
	writer.Flush()
}
```
