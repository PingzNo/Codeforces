# 50A. Domino piling

## Problem

You are given a rectangular board of $M \times N$ squares. Also you are given an unlimited number of standard domino pieces of $2 \times 1$ squares. You are allowed to rotate the pieces. You are asked to place as many dominoes as possible on the board so as to meet the following conditions:

  - 1. Each domino completely covers two squares.
  - 2. No two dominoes overlap.
  - 3. Each domino lies entirely inside the board. It is allowed to touch the edges of the board.

Find the maximum number of dominoes, which can be placed under these restrictions.

[Link to CodeForces](https://codeforces.com/problemset/problem/50/A)

## Input

In a single line you are given two integers $M$ and $N$ — board sizes in squares ($1 \leq M \leq N \leq 16$).

## Output

Output one number — the maximal number of dominoes, which can be placed.

## Examples

### Input 1

```
2 4
```

### Output 1

```
4
```

### Input 2

```
3 3
```

### Output 2

```
4
```

## Analysis

The approach to solving this problem is similar to the one of [1A Theatre Square](https://blog.zhao.no/day-4-1a-theatre-square). The tricky part is that one
domino must be placed entirely inside the board, according to the problem statement. As a result, the dominos on the last row or column with odd board height or
width should be rotated to meet this requirement. This part is different from the problem [1A Theatre Square](https://blog.zhao.no/day-4-1a-theatre-square).

## Solutions

### Constraints

  - Time limit per test: 2 seconds
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

### GNU C++17 7.3.0

| Problem |    Lang   |  Verdict | Time  | Memory |
|:-------:|:---------:|:--------:|:-----:|:------:|
|  50A-14 | GNU C++17 | Accepted | 30 ms |  8 KB  |

[Link to source code](solution.cpp)

```c++
#include <iostream>


void solution() {
    int height, width;
    std::cin >> height >> width;

    int half_height = height / 2;
    int half_width = width / 2;
    int result = 2 * half_height * half_width;
    if ((height % 2 == 1) && (width % 2 == 1)) {
        result += half_height + half_width;
    } else if ((height % 2 == 1) && (width % 2 == 0)) {
        result += half_width;
    } else if ((height % 2 == 0) && (width % 2 == 1)) {
        result += half_height;
    }

    std::cout << result << '\n';
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

### C# Mono 6.8

| Problem |    Lang   |  Verdict | Time   |   Memory  |
|:-------:|:---------:|:--------:|:------:|:---------:|
|  50A-14 |   Mono C# | Accepted | 154 ms |  3612 KB  |

[Link to source code](solution.cs)

```c#
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;


namespace Codeforces
{
    class Program
    {
        static void Solution()
        {
            List<int> line = System.Console.ReadLine().Split().Select(int.Parse).ToList();
            int height = line[0], width = line[1];

            int half_height = height / 2;
            int half_width = width / 2;
            int result = 2 * half_height * half_width;
            if ((height % 2 == 1) && (width % 2 == 1))
            {
                result += half_height + half_width;
            }
            else if ((height % 2 == 1) && (width % 2 == 0))
            {
                result += half_width;
            }
            else if ((height % 2 == 0) && (width % 2 == 1))
            {
                result += half_height;
            }

            System.Console.WriteLine("{0}", result);
        }

        static void Main(string[] args)
        {
            Solution();
        }
    }
}
```

### Python 3.8.10

| Problem |    Lang   |  Verdict | Time   |   Memory  |
|:-------:|:---------:|:--------:|:------:|:---------:|
|  50A-14 |     Go    | Accepted | 92 ms  |    0 KB   |

[Link to source code](solution.py)

```python
def solution():
    height, width = [int(x) for x in input().split()]

    half_height = height // 2
    half_width = width // 2

    result = 2 * half_height * half_width
    if height % 2 == 1 and width % 2 == 1:
        result += half_height + half_width

    elif height % 2 == 1 and width % 2 == 0:
        result += half_width

    elif height % 2 == 0 and width % 2 == 1:
        result += half_height

    print(result)


if __name__ == "__main__":
    solution()
```

### Go 1.17.5

| Problem |    Lang   |  Verdict | Time   |   Memory  |
|:-------:|:---------:|:--------:|:------:|:---------:|
|  50A-14 |     Go    | Accepted | 30 ms  |   64 KB   |

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

func Printf(format string, x ...interface{}) {
	fmt.Fprintf(writer, format, x...)
}

func Scanf(format string, x ...interface{}) {
	fmt.Fscanf(reader, format, x...)
}

func Solution() {
	var height, width int
	Scanf("%d %d\n", &height, &width)

	half_height := height / 2
	half_width := width / 2

	var result int = 2 * half_height * half_width
	if (height%2 == 1) && (width%2 == 1) {
		result += half_height + half_width
	} else if (height%2 == 1) && (width%2 == 0) {
		result += half_width
	} else if (height%2 == 0) && (width%2 == 1) {
		result += half_height
	}

	Printf("%d\n", result)
}

func main() {
	Solution()
	writer.Flush()
}
```
