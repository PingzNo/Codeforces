# 1A. Theatre Square

## Problem

Theatre Square in the capital city of Berland has a rectangular shape with the size $n \times m$ meters. On the occasion of the city's anniversary, a decision was taken to pave the Square with square granite flagstones. Each flagstone is of the size $a \times a$.

What is the least number of flagstones needed to pave the Square? It's allowed to cover the surface larger than the Theatre Square, but the Square has to be covered. It's not allowed to break the flagstones. The sides of flagstones should be parallel to the sides of the Square.

[Link to CodeForces](https://codeforces.com/problemset/problem/1/A)

## Input

The input contains three positive integer numbers in the first line: $n$, $m$ and $a$ ($1 \leq  n$, $m$, $a \leq 109$).

## Output

Write the needed number of flagstones.

## Examples

### Input

```
6 6 4
```

### Output

```
4
```

## Analysis

The key to solving this problem is checking if the square height or width is an even number or an odd one. As it is illustrated in the figure below, if
the square height or width is an even number, then one can place the flagstone to precisely fit the square along that axis (solid lines). In the case that the square height or
width is an odd number; exactly half of the last flagstones will be outside the square along the axis (dashed lines). However, no matter which case it is, the
total number of flagstones used will be equal to the one in the case of even height and width.

![Illustration](1A.png)

## Solutions

### Constraints

  - Time limit per test: 1 second
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

### Python 3.8.10

| Problem |    Lang   |  Verdict | Time  | Memory |
|:-------:|:---------:|:--------:|:-----:|:------:|
|   1A-10 |  Python 3 | Accepted | 46 ms |  0 KB  |

[Link to source code](solution.py)

```python
def solution():
	height, width, size = [int(x) for x in input().split()]
	height_ratio = height // size if height % size == 0 else height // size + 1
	width_ratio = width // size if width % size == 0 else width // size + 1
	area = height_ratio * width_ratio
	print(area)


if __name__ == "__main__":
	solution()
```

### GNU C++17 7.3.0

| Problem |    Lang   |  Verdict | Time  | Memory |
|:-------:|:---------:|:--------:|:-----:|:------:|
|   1A-10 | GNU C++17 | Accepted | 15 ms |  4 KB  |

[Link to source code](solution.cpp)

```c++
#include <iostream>


void solution() {
    unsigned long long height, width, size;
    std::cin >> height >> width >> size;

    unsigned long long height_ratio = (height % size == 0) ? (height / size) : ((height / size) + 1);
    unsigned long long width_ratio = (width % size == 0) ? (width / size) : ((width / size) + 1);
    std::cout << height_ratio * width_ratio << '\n';
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

### C# 10, .NET SDK 6.0

| Problem |    Lang   |  Verdict | Time  | Memory |
|:-------:|:---------:|:--------:|:-----:|:------:|
|   1A-10 |   C# 10   | Accepted | 46 ms |  0 KB  |

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
            List<UInt64> line = System.Console.ReadLine().Split().Select(UInt64.Parse).ToList();
            UInt64 height = line[0], width = line[1], size = line[2];

            UInt64 height_ratio = (height % size == 0) ? (height / size) : ((height / size) + 1);
            UInt64 width_ratio = (width % size == 0) ? (width / size) : ((width / size) + 1);

            System.Console.WriteLine("{0}", height_ratio * width_ratio);
        }

        static void Main(string[] args)
        {
            Solution();
        }
    }
}
```

### Go 1.17.5

| Problem |    Lang   |  Verdict | Time  | Memory |
|:-------:|:---------:|:--------:|:-----:|:------:|
|   1A-10 |     Go    | Accepted | 30 ms | 60 KB  |

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

func Printf(f string, a ...interface{}) {
	fmt.Fprintf(writer, f, a...)
}

func Scanf(f string, a ...interface{}) {
	fmt.Fscanf(reader, f, a...)
}

func Solution() {
	var height, width, size uint64
	Scanf("%d %d %d\n", &height, &width, &size)

	var heightRatio uint64
	if height%size == 0 {
		heightRatio = height / size
	} else {
		heightRatio = height/size + 1
	}

	var widthRatio uint64
	if width%size == 0 {
		widthRatio = width / size
	} else {
		widthRatio = width/size + 1
	}

	Printf("%d\n", heightRatio*widthRatio)
}

func main() {
	defer writer.Flush()
	Solution()
}
```
