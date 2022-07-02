# 41A. Translation

## Problem

The translation from the Berland language into the Birland language is not an easy task. Those languages are very similar: a berlandish word differs from a birlandish word with the same meaning a little: it is spelled (and pronounced) reversely. For example, a Berlandish word code corresponds to a Birlandish word edoc. However, it's easy to make a mistake during the «translation». Vasya translated word s from Berlandish into Birlandish as t. Help him: find out if he translated the word correctly.

[Link to CodeForces](https://codeforces.com/problemset/problem/1/A)

## Input

The first line contains word $s$, the second line contains word $t$. The words consist of lowercase Latin letters. The input data do not consist unnecessary spaces.
The words are not empty and their lengths do not exceed 100 symbols.

## Output

If the word $t$ is a word $s$, written reversely, print YES, otherwise print NO.

## Examples

### Input

```
code
edoc
```

### Output

```
YES
```

### Input

```
abb
aba
```

### Output

```
NO
```

### Input

```
code
code
```

### Output

```
NO
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
| 41A - 9 |    Go     | Accepted | 62 ms |  60 KB |

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
	var berlandish, birlandish string
	Fscanf("%s\n%s\n", &berlandish, &birlandish)

	if len(berlandish) != len(birlandish) {
		Fprintf("NO\n")
		return
	}

	for i := 0; i < len(berlandish); i++ {
		if berlandish[i] != birlandish[len(birlandish)-i-1] {
			Fprintf("NO\n")
			return
		}
	}

	Fprintf("YES\n")
}

func main() {
	Solution()
	writer.Flush()
}
```
