# 118A. String Task

## Problem

Petya started to attend programming lessons. On the first lesson his task was to write a simple program. The program was supposed to do the following: in the given string, consisting if uppercase and lowercase Latin letters, it:

  - Deletes all the vowels,
  - Inserts a character "." before each consonant,
  - Replaces all uppercase consonants with corresponding lowercase ones.

Vowels are letters "A", "O", "Y", "E", "U", "I", and the rest are consonants. The program's input is exactly one string, it should return the output as a single string, resulting after the program's processing the initial string.

Help Petya cope with this easy task.

[Link to CodeForces](https://codeforces.com/problemset/problem/118/A)

## Input

The first line represents input string of Petya's program. This string only consists of uppercase and lowercase Latin letters and its length is from 1 to 100,
inclusive.

## Output

Print the resulting string. It is guaranteed that this string is not empty.

## Examples

### Input 1

[Link to file](input_0.txt)

```
tour
```

### Output 1

[Link to file](expected_0.txt)

```
.t.r
```

### Input 2

[Link to file](input_1.txt)

```
Codeforces
```

### Output 2

[Link to file](expected_1.txt)

```
.c.d.f.r.c.s
```

### Input 3

[Link to file](input_2.txt)

```
aBAcAba
```

### Output 3

[Link to file](expected_2.txt)

```
.b.c.b
```

## Solutions

### Constraints

  - Time limit per test: 2 seconds
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

### GNU C++17 7.3.0

|  Problem  |    Lang   |  Verdict |  Time | Memory |
|:---------:|:---------:|:--------:|:-----:|:------:|
| 118A - 13 | GNU C++17 | Accepted | 30 ms |  8 KB  |

[Link to source code](solution.cpp)

```c++
#include <iostream>
#include <string>
#include <cctype>


void solution() {
    std::string text;
    std::cin >> text;

    for (std::string::const_iterator p = text.begin(); p != text.end(); ++p) {
        switch (tolower(*p)) {
            case 'a':
            case 'o':
            case 'y':
            case 'e':
            case 'u':
            case 'i':
                continue;

            default:
                std::cout << '.' << char(tolower(*p));
                break;
        }
    }

    std::cout << '\n';
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

|  Problem  |    Lang   |  Verdict |  Time | Memory |
|:---------:|:---------:|:--------:|:-----:|:------:|
| 118A - 13 |    Go     | Accepted | 30 ms | 76 KB  |

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
	var text string
	Scanf("%s\n", &text)
	text = strings.ToLower(text)

	for i := 0; i < len(text); i++ {
		switch text[i] {
		case 'a':
		case 'o':
		case 'y':
		case 'e':
		case 'u':
		case 'i':
			continue

		default:
			Printf(".%c", text[i])
			break
		}
	}

	Printf("\n")
}

func main() {
	Solution()
	writer.Flush()
}
```

### Python 3.8.10

|  Problem  |    Lang   |  Verdict |  Time | Memory |
|:---------:|:---------:|:--------:|:-----:|:------:|
| 118A - 13 |  Python   | Accepted | 92 ms |  0 KB  |

[Link to source code](solution.py)

```python
def solution():
    for c in input().lower():
        if c == 'a' or c == 'o' or c == 'y' or c == 'e' or c == 'u' or c == 'i':
            continue

        print(f".{c}", end="")

    print()


if __name__ == "__main__":
    solution()
```
