# 112A. Petya and Strings

## Problem

Little Petya loves presents. His mum bought him two strings of the same size for his birthday. The strings consist of uppercase and lowercase Latin letters. Now
Petya wants to compare those two strings lexicographically. The letters' case does not matter, that is an uppercase letter is considered equivalent to the
corresponding lowercase letter. Help Petya perform the comparison.

[Link to CodeForces](https://codeforces.com/problemset/problem/112/A)

## Input

Each of the first two lines contains a bought string. The strings' lengths range from 1 to 100 inclusive. It is guaranteed that the strings are of the same length and also consist of uppercase and lowercase Latin letters.

## Output

If the first string is less than the second one, print "-1". If the second string is less than the first one, print "1". If the strings are equal, print "0". Note that the letters' case is not taken into consideration when the strings are compared.

## Examples

### Input 1

[Link to file](https://github.com/PingzNo/Codeforces/blob/master/0112A.Petya.and.Strings/input_0.txt)

```
aaaa
aaaA
```

### Output 1

[Link to file](https://github.com/PingzNo/Codeforces/blob/master/0112A.Petya.and.Strings/expected_0.txt)

```
0
```

### Input 2

[Link to file](https://github.com/PingzNo/Codeforces/blob/master/0112A.Petya.and.Strings/input_1.txt)

```
abs
Abz
```

### Output 2

[Link to file](https://github.com/PingzNo/Codeforces/blob/master/0112A.Petya.and.Strings/expected_1.txt)

```
-1
```

### Input 3

[Link to file](https://github.com/PingzNo/Codeforces/blob/master/0112A.Petya.and.Strings/input_2.txt)

```
abcdefg
AbCdEfF
```

### Output 3

[Link to file](https://github.com/PingzNo/Codeforces/blob/master/0112A.Petya.and.Strings/expected_2.txt)

```
1
```

## Note

If you want more formal information about the lexicographical order (also known as the "dictionary order" or "alphabetical order"), you can visit the following site:

[http://en.wikipedia.org/wiki/Lexicographical_order](http://en.wikipedia.org/wiki/Lexicographical_order)

## Analysis

The solution to this problem is a straightforward implementation of a string comparison operator by ignoring the character cases.

## Solutions

### Constraints

  - Time limit per test: 2 seconds
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

### GNU C++17 7.3.0

|  Problem  |    Lang   |  Verdict |  Time | Memory |
|:---------:|:---------:|:--------:|:-----:|:------:|
| 112A - 11 | GNU C++17 | Accepted | 30 ms |  8 KB  |

[Link to source code](solution.cpp)

```c++
#include <iostream>
#include <string>
#include <cctype>


void solution() {
    std::string first, second;
    std::cin >> first >> second;

    int result = 0;
    std::string::const_iterator p = first.begin();
    std::string::const_iterator q = second.begin();
    while (p != first.end()) {
        if (toupper(*p) < toupper(*q)) {
            result = -1;
            break;
        }

        if (toupper(*p) > toupper(*q)) {
            result = 1;
            break;
        }

        ++p;
        ++q;
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

### Python 3.8.10

|  Problem  |    Lang   |  Verdict |  Time | Memory |
|:---------:|:---------:|:--------:|:-----:|:------:|
| 112A - 11 |   Python  | Accepted | 92 ms |  0 KB  |

[Link to source code](solution.py)

```python
def solution():
    first = input().upper()
    second = input().upper()

    result = 0
    for i, c in enumerate(first):
        if c < second[i]:
            result = -1
            break

        if c > second[i]:
            result = 1
            break

    print(result)


if __name__ == "__main__":
    solution()
```

### C# Mono 6.8

|  Problem  |    Lang   |  Verdict |  Time | Memory |
|:---------:|:---------:|:--------:|:-----:|:------:|
| 112A - 11 |  Mono C#  | Accepted | 92 ms |  0 KB  |

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
            String first = System.Console.ReadLine().ToUpper();
            String second = System.Console.ReadLine().ToUpper();

            int result = 0;
            for (int i = 0; i < first.Length; ++i)
            {
                if (first[i] < second[i])
                {
                    result = -1;
                    break;
                }

                if (first[i] > second[i])
                {
                    result = 1;
                    break;
                }
            }

            System.Console.WriteLine(result);
        }

        static void Main(string[] args)
        {
            Solution();
        }
    }
}
```

### Go 1.17.5

|  Problem  |    Lang   |  Verdict |  Time | Memory |
|:---------:|:---------:|:--------:|:-----:|:------:|
| 112A - 11 |    Go     | Accepted | 62 ms | 68 KB  |

[Link to source code](solution.go)

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	var first, second string
	Scanf("%s\n%s\n", &first, &second)
	first = strings.ToUpper(first)
	second = strings.ToUpper(second)

	result := 0
	for i := 0; i < len(first); i++ {
		if first[i] < second[i] {
			result = -1
			break
		}

		if first[i] > second[i] {
			result = 1
			break
		}
	}

	Printf("%d\n", result)
}

func main() {
	Solution()
	writer.Flush()
}
```
