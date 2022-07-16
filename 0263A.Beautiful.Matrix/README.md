# 263A. Beautiful Matrix

## Problem

You've got a $5 \times 5$ matrix, consisting of 24 zeroes and a single number one. Let's index the matrix rows by numbers from 1 to 5 from top to bottom, let's index the matrix columns by numbers from 1 to 5 from left to right. In one move, you are allowed to apply one of the two following transformations to the matrix:

Swap two neighboring matrix rows, that is, rows with indexes $i$ and $i + 1$ for some integer $i$ ($1 \leq i \lt 5$).
Swap two neighboring matrix columns, that is, columns with indexes $j$ and $j + 1$ for some integer $j$ ($1 \leq j \lt 5$).
You think that a matrix looks beautiful, if the single number one of the matrix is located in its middle (in the cell that is on the intersection of the third row and the third column). Count the minimum number of moves needed to make the matrix beautiful.

[Link to CodeForces](https://codeforces.com/problemset/problem/263/A)

## Input

The input consists of five lines, each line contains five integers: the $j$-th integer in the $i$-th line of the input represents the element of the matrix that is
located on the intersection of the $i$-th row and the $j$-th column. It is guaranteed that the matrix consists of 24 zeroes and a single number one.

## Output

Print a single integer — the minimum number of moves needed to make the matrix beautiful.

## Examples

### Input

```
0 0 0 0 0
0 0 0 0 1
0 0 0 0 0
0 0 0 0 0
0 0 0 0 0
```

### Output

```
3
```

### Input

```
0 0 0 0 0
0 0 0 0 0
0 1 0 0 0
0 0 0 0 0
0 0 0 0 0
```

### Output

```
1
```

## Constraints

  - Time limit per test: 2 seconds
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

## Solutions

### GNU C++17 7.3.0

| Problem  |    Lang   |  Verdict | Time  | Memory |
|:--------:|:---------:|:--------:|:-----:|:------:|
| 263A - 8 | GNU C++17 | Accepted | 30 ms |  4 KB  |

[Link to source code](solution.cpp)

```c++
#include <iostream>


void solution() {
    int number = 0;
    for (int row_index = 0; row_index < 5; ++row_index) {
        for (int col_index = 0; col_index < 5; ++col_index) {
            std::cin >> number;
            if (number == 1) {
                std::cout << abs(row_index - 2) + abs(col_index - 2) << '\n';
                return;
            }
        }
    }
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

### Go 1.17.5

| Problem  |    Lang   |  Verdict | Time  | Memory |
|:--------:|:---------:|:--------:|:-----:|:------:|
| 263A - 8 |    GO     | Accepted | 60 ms |  0 KB  |

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
	var number int
	for rowIndex := 0; rowIndex < 5; rowIndex++ {
		for colIndex := 0; colIndex < 5; colIndex++ {
			Fscanf("%d", &number)
			if number == 1 {
				rowDistance := int(math.Abs(float64(rowIndex - 2)))
				colDistance := int(math.Abs(float64(colIndex - 2)))
				Fprintf("%d\n", rowDistance+colDistance)
				return
			}
		}
		Fscanf("\n")
	}
}

func main() {
	Solution()
	writer.Flush()
}
```
