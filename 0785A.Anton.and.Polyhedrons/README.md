# 785A. Anton and Polyhedrons

## Problem

Anton's favourite geometric figures are regular polyhedrons. Note that there are five kinds of regular polyhedrons:

Tetrahedron. Tetrahedron has 4 triangular faces.
Cube. Cube has 6 square faces.
Octahedron. Octahedron has 8 triangular faces.
Dodecahedron. Dodecahedron has 12 pentagonal faces.
Icosahedron. Icosahedron has 20 triangular faces.
All five kinds of polyhedrons are shown on the picture below:

![Illustration](illustration.png)

Anton has a collection of $n$ polyhedrons. One day he decided to know, how many faces his polyhedrons have in total. Help Anton and find this number!

[Link to CodeForces](https://codeforces.com/problemset/problem/785/A)

## Input

The first line of the input contains a single integer $n$ ($1 \leq n \leq 200 000$) — the number of polyhedrons in Anton's collection.

Each of the following $n$ lines of the input contains a string $s_i$ — the name of the $i$-th polyhedron in Anton's collection. The string can look like this:

  - "Tetrahedron" (without quotes), if the $i$-th polyhedron in Anton's collection is a tetrahedron.
  - "Cube" (without quotes), if the $i$-th polyhedron in Anton's collection is a cube.
  - "Octahedron" (without quotes), if the $i$-th polyhedron in Anton's collection is an octahedron.
  - "Dodecahedron" (without quotes), if the $i$-th polyhedron in Anton's collection is a dodecahedron.
  - "Icosahedron" (without quotes), if the $i$-th polyhedron in Anton's collection is an icosahedron.

## Output

Output one number — the total number of faces in all the polyhedrons in Anton's collection.

## Examples

### Input

```
4
Icosahedron
Cube
Tetrahedron
Dodecahedron
```

### Output

```
42
```

### Input

```
3
Dodecahedron
Octahedron
Octahedron
```

### Output

```
28
```

## Note

In the first sample Anton has one icosahedron, one cube, one tetrahedron and one dodecahedron. Icosahedron has 20 faces, cube has 6 faces, tetrahedron has 4 faces and dodecahedron has 12 faces. In total, they have 20 + 6 + 4 + 12 = 42 faces.

## Solutions

### Constraints

  - Time limit per test: 2 seconds
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

### Go 1.17.5

|  Problem  |    Lang   |  Verdict | Time   |   Memory  |
|:---------:|:---------:|:--------:|:------:|:---------:|
| 785A - 19 |    Go     | Accepted | 124 ms |  3184 KB  |

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
	var cubeCount int
	Fscanf("%d\n", &cubeCount)

	faceCount := 0

	var cubeName string
	for cubeIndex := 0; cubeIndex < cubeCount; cubeIndex++ {
		Fscanf("%s\n", &cubeName)
		switch cubeName[0] {
		case 'T':
			faceCount += 4
		case 'C':
			faceCount += 6
		case 'O':
			faceCount += 8
		case 'D':
			faceCount += 12
		case 'I':
			faceCount += 20
		}
	}

	Fprintf("%d\n", faceCount)
}

func main() {
	Solution()
	writer.Flush()
}
```
