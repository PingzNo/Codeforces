# 281A. Word Capitalization

## Problem

Capitalization is writing a word with its first letter as a capital letter. Your task is to capitalize the given word.

Note, that during capitalization all the letters except the first one remains unchanged.

[Link to CodeForces](https://codeforces.com/problemset/problem/281/A)

## Input

A single line contains a non-empty word. This word consists of lowercase and uppercase English letters. The length of the word will not exceed $10^3$.

## Output

Output the given word after capitalization.

## Examples

### Input 1

[Link to file](input_0.txt)

```
ApPLe
```

### Output 1

[Link to file](expected_0.txt)

```
ApPLe
```

### Input 2

[Link to file](input_1.txt)

```
konjac
```

### Output 2

[Link to file](expected_1.txt)

```
Konjac
```

## Solutions

### Constraints

  - Time limit per test: 2 seconds
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

### Python 3.8.10

| Problem |    Lang   |  Verdict | Time  | Memory |
|:-------:|:---------:|:--------:|:-----:|:------:|
| 281A-22 |   Python  | Accepted | 62 ms |  0 KB  |

[Link to source code](solution.py)

```python
def solution():
    word = input()
    print(f"{word[0].upper()}{word[1:]}")


if __name__ == "__main__":
    solution()
```

### Go 1.17.5

| Problem |    Lang   |  Verdict | Time  | Memory |
|:-------:|:---------:|:--------:|:-----:|:------:|
| 281A-22 |     Go    | Accepted | 30 ms | 52 KB  |

[Link to source code](solution.go)

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
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
	var word string
	Fscanf("%s\n", &word)
	Fprintf("%c%s\n", unicode.ToUpper(rune(word[0])), word[1:])
}

func main() {
	Solution()
	writer.Flush()
}
```
