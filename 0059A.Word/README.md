# 59A. Word

## Problem

Vasya is very upset that many people on the Net mix uppercase and lowercase letters in one word. That's why he decided to invent an extension for his favorite browser that would change the letters' register in every word so that it either only consisted of lowercase letters or, vice versa, only of uppercase ones. At that as little as possible letters should be changed in the word. For example, the word HoUse must be replaced with house, and the word ViP — with VIP. If a word contains an equal number of uppercase and lowercase letters, you should replace all the letters with lowercase ones. For example, maTRIx should be replaced by matrix. Your task is to use the given method on one given word.

[Link to CodeForces](https://codeforces.com/problemset/problem/59/A)

## Input

The first line contains $a$ word $s$ — it consists of uppercase and lowercase Latin letters and possesses the length from 1 to 100.

## Output

Print the corrected word s. If the given word s has strictly more uppercase letters, make the word written in the uppercase register, otherwise - in the lowercase one.

## Examples

### Input

```
HoUse
```

### Output

```
house
```

### Input

```
ViP
```

### Output

```
VIP
```

### Input

```
maTRIx
```

### Output

```
matrix
```

## Solutions

### Constraints

  - Time limit per test: 2 seconds
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

### Go 1.17.5

| Problem |    Lang   |  Verdict | Time  | Memory |
|:-------:|:---------:|:--------:|:-----:|:------:|
|  59A-5  |    Go     | Accepted | 60 ms | 20 KB  |

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
	var word string
	Fscanf("%s\n", &word)

	var lowerCaseCount, upperCaseCount int
	for _, c := range word {
		if c >= 'a' && c <= 'z' {
			lowerCaseCount++
		} else if c >= 'A' && c <= 'Z' {
			upperCaseCount++
		}
	}

	if lowerCaseCount < upperCaseCount {
		Fprintf("%s\n", strings.ToUpper(word))
	} else {
		Fprintf("%s\n", strings.ToLower(word))
	}
}

func main() {
	Solution()
	writer.Flush()
}
```
