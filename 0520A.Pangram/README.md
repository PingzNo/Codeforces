# 520A. Pangram

## Problem

A word or a sentence in some language is called a pangram if all the characters of the alphabet of this language appear in it at least once. Pangrams are often used to demonstrate fonts in printing or test the output devices.

You are given a string consisting of lowercase and uppercase Latin letters. Check whether this string is a pangram. We say that the string contains a letter of the Latin alphabet if this letter occurs in the string in uppercase or lowercase.

[Link to CodeForces](https://codeforces.com/problemset/problem/520/A)

## Input

The first line contains a single integer $n$ ($1 \leq n \leq 100$) — the number of characters in the string.

The second line contains the string. The string consists only of uppercase and lowercase Latin letters.

## Output

Output "YES", if the string is a pangram and "NO" otherwise.

## Examples

### Input

```
12
toosmallword
```

### Output

```
NO
```

### Input

```
35
TheQuickBrownFoxJumpsOverTheLazyDog
```

### Output

```
YES
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
| 520A - 24 |    Go     | Accepted | 31 ms | 76 KB  |

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

func Fprintf(format string, x ...interface{}) {
	fmt.Fprintf(writer, format, x...)
}

func Fscanf(format string, x ...interface{}) {
	fmt.Fscanf(reader, format, x...)
}

func Solution() {
	var letterCount int
	var letters string
	Fscanf("%d\n%s\n", &letterCount, &letters)

	letterMap := make(map[rune]bool)
	for _, letter := range strings.ToUpper(letters) {
		letterMap[letter] = true
	}

	if len(letterMap) == 26 {
		Fprintf("YES\n")
		return
	}

	Fprintf("NO\n")
}

func main() {
	Solution()
	writer.Flush()
}
```
