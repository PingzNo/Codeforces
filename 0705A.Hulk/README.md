# 705A. Hulk

## Problem

Dr. Bruce Banner hates his enemies (like others don't). As we all know, he can barely talk when he turns into the incredible Hulk. That's why he asked you to help him to express his feelings.

Hulk likes the Inception so much, and like that his feelings are complicated. They have n layers. The first layer is hate, second one is love, third one is hate and so on...

For example if $n = 1$, then his feeling is "I hate it" or if $n = 2$ it's "I hate that I love it", and if $n = 3$ it's "I hate that I love that I hate it" and so on.

Please help Dr. Banner.

[Link to CodeForces](https://codeforces.com/problemset/problem/705/A)

## Input

The only line of the input contains a single integer $n$ ($1 \leq n \leq 100$) — the number of layers of love and hate.

## Output

Print Dr.Banner's feeling in one line.

## Examples

### Input

```
1
```

### Output

```
I hate it
```

### Input

```
2
```

### Output

```
I hate that I love it
```

### Input

```
3
```

### Output

```
I hate that I love that I hate it
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
| 705A - 8 |    Go     | Accepted | 15 ms | 16 KB  |

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
	var layerCount int
	Fscanf("%d\n", &layerCount)

	actions := []string{"hate", "love"}
	actionIndex := 0

	for layerIndex := 1; layerIndex <= layerCount; layerIndex++ {
		if layerIndex == layerCount {
			Fprintf("I %s it\n", actions[actionIndex])
		} else {
			Fprintf("I %s that ", actions[actionIndex])
		}

		actionIndex = (actionIndex + 1) % 2
	}
}

func main() {
	Solution()
	writer.Flush()
}
```
