# 443A. Anton and Letters

## Problem

Recently, Anton has found a set. The set consists of small English letters. Anton carefully wrote out all the letters from the set in one line, separated by a comma. He also added an opening curved bracket at the beginning of the line and a closing curved bracket at the end of the line.

Unfortunately, from time to time Anton would forget writing some letter and write it again. He asks you to count the total number of distinct letters in his set.

[Link to CodeForces](https://codeforces.com/problemset/problem/443/A)

## Input

The first and the single line contains the set of letters. The length of the line doesn't exceed 1000. It is guaranteed that the line starts from an opening curved bracket and ends with a closing curved bracket. Between them, small English letters are listed, separated by a comma. Each comma is followed by a space.

## Output

Print a single number — the number of distinct letters in Anton's set.

## Examples

### Input

```
{a, b, c}
```

### Output

```
3
```

### Input

```
{b, a, b, a}
```

### Output

```
2
```

### Input

```
{}
```

### Output

```
0
```

## Solutions

### Constraints

  - Time limit per test: 2 seconds
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

### Go 1.17.5

|  Problem  |    Lang   |  Verdict | Time  | Memory |
|:---------:|:---------:|:--------:|:-----:|:------:|
| 443A - 14 |     Go    | Accepted | 31 ms | 28 KB  |

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
	var line string
	line, _ = reader.ReadString('\n')

	letters := map[rune]int{}
	for _, c := range line {
		if unicode.IsLower(c) {
			letters[c]++
		}
	}

	Fprintf("%d\n", len(letters))
}

func main() {
	Solution()
	writer.Flush()
}
```
