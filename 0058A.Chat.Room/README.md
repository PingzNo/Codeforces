# 58A. Chat room

## Problem

Vasya has recently learned to type and log on to the Internet. He immediately entered a chat room and decided to say hello to everybody. Vasya typed the word s.
It is considered that Vasya managed to say hello if several letters can be deleted from the typed word so that it resulted in the word "hello". For example, if
Vasya types the word "ahhellllloou", it will be considered that he said hello, and if he types "hlelo", it will be considered that Vasya got misunderstood and
he didn't manage to say hello. Determine whether Vasya managed to say hello by the given word s.

[Link to CodeForces](https://codeforces.com/problemset/problem/1/A)

## Input

The first and only line contains the word $s$, which Vasya typed. This word consisits of small Latin letters, its length is no less that 1 and no more than 100
letters.

## Output

If Vasya managed to say hello, print "YES", otherwise print "NO".

## Examples

### Input

```
ahhellllloou
```

### Output

```
YES
```

### Input

```
hlelo
```

### Output

```
NO
```

## Solutions

### Constraints

  - Time limit per test: 1 second
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

### Go 1.17.5

| Problem  |    Lang   |  Verdict | Time  | Memory |
|:--------:|:---------:|:--------:|:-----:|:------:|
| 58A - 10 |    Go     | Accepted | 31 ms |  56 KB  |

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

func Fprintf(format string, x ...interface{}) {
	fmt.Fprintf(writer, format, x...)
}

func Fscanf(format string, x ...interface{}) {
	fmt.Fscanf(reader, format, x...)
}

func Solution() {
	var word string
	Fscanf("%s\n", &word)
	reference := "hello"

	refIndex := 0
	wordIndex := 0
	for wordIndex < len(word) && refIndex < len(reference) {
		if word[wordIndex] == reference[refIndex] {
			refIndex++
		}
		wordIndex++
	}

	if refIndex == len(reference) {
		Fprintf("YES\n")
	} else {
		Fprintf("NO\n")
	}
}

func main() {
	Solution()
	writer.Flush()
}
```
