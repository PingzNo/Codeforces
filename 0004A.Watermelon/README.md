# 4A. Watermelon

## Problem

One hot summer day Pete and his friend Billy decided to buy a watermelon. They chose the biggest and the ripest one, in their opinion. After that the watermelon was weighed, and the scales showed w kilos. They rushed home, dying of thirst, and decided to divide the berry, however they faced a hard problem.

Pete and Billy are great fans of even numbers, that's why they want to divide the watermelon in such a way that each of the two parts weighs even number of
kilos, at the same time it is not obligatory that the parts are equal. The boys are extremely tired and want to start their meal as soon as possible, that's why
you should help them and find out, if they can divide the watermelon in the way they want. For sure, each of them should get a part of positive weight.

[Link to CodeForces](https://codeforces.com/problemset/problem/4/A)

## Input

The first (and the only) input line contains integer number $w$ ($1 \leq w \leq 100$) — the weight of the watermelon bought by the boys.

## Output

Print YES, if the boys can divide the watermelon into two parts, each of them weighing even number of kilos; and NO in the opposite case.

## Examples

### Input

```
8
```

### Output

```
YES
```

## Note

For example, the boys can divide the watermelon into two parts of 2 and 6 kilos respectively (another variant — two parts of 4 and 4 kilos).

## Constraints

  - Time limit per test: 1 second
  - Memory limit per test: 64 megabytes
  - Input: standard input
  - Output: standard output

## Solutions

### GNU C++17 7.3.0

| Problem |    Lang   |  Verdict | Time | Memory |
|:-------:|:---------:|:--------:|:----:|:------:|
|   4A-8  | GNU C++17 | Accepted | 0 ms |  4 KB  |

[Link to source code](solution.cpp)

```c++
#include <iostream>


void solution() {
    int weight = 0;
    std::cin >> weight;
    std::cout << ((weight > 2 && weight % 2 == 0) ? "YES" : "NO") << '\n';
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

| Problem |    Lang   |  Verdict | Time  | Memory |
|:-------:|:---------:|:--------:|:-----:|:------:|
|   4A-8  |  Mono C#  | Accepted | 92 ms |  0 KB  |

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
            int w = Convert.ToInt32(System.Console.ReadLine());
            System.Console.WriteLine("{0}", ((w > 2 && w % 2 == 0) ? "YES" : "NO"));
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
|   4A-8  |     Go    | Accepted | 30 ms | 56 KB  |

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
	var weight int
	Scanf("%d\n", &weight)
	if weight > 2 && weight%2 == 0 {
		Printf("YES\n")
	} else {
		Printf("NO\n")
	}
}

func main() {
	Solution()
	writer.Flush()
}
```
