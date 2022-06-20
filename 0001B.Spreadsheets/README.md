# 1B. Spreadsheets

## Problem

In the popular spreadsheets systems (for example, in Excel) the following numeration of columns is used. The first column has number A, the second — number B, etc. till column 26 that is marked by Z. Then there are two-letter numbers: column 27 has number AA, 28 — AB, column 52 is marked by AZ. After ZZ there follow three-letter numbers, etc.

The rows are marked by integer numbers starting with 1. The cell name is the concatenation of the column and the row numbers. For example, BC23 is the name for the cell that is in column 55, row 23.

Sometimes another numeration system is used: RXCY, where X and Y are integer numbers, showing the column and the row numbers respectfully. For instance, R23C55 is the cell from the previous example.

Your task is to write a program that reads the given sequence of cell coordinates and produce each item written according to the rules of another numeration system.

[Link to CodeForces](https://codeforces.com/problemset/problem/1/B)

## Input

The first line of the input contains integer number $n$ ($1 \leq n \leq 10^5$), the number of coordinates in the test. Then there follow $n$ lines, each of them contains coordinates. All the coordinates are correct, there are no cells with the column and/or the row numbers larger than $10^6$.

## Output

Write $n$ lines, each line should contain a cell coordinates in the other numeration system.

## Examples

### Input

```
2
R23C55
BC23
```

### Output

```
BC23
R23C55
```

## Solutions

### Constraints

  - Time limit per test: 10 seconds
  - Memory limit per test: 64 megabytes
  - Input: standard input
  - Output: standard output

### Go 1.17.5

| Problem |    Lang   |  Verdict |  Time  |  Memory  |
|:-------:|:---------:|:--------:|:------:|:--------:|
| 1B - 9  |    Go     | Accepted | 248 ms | 4308 KB  |

[Link to source code](solution.go)

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	var caseCount int
	Fscanf("%d\n", &caseCount)

	for caseIndex := 0; caseIndex < caseCount; caseIndex++ {
		var text string
		Fscanf("%s\n", &text)

		var pos int
		for pos < len(text) && text[pos] != 'C' {
			pos++
		}

		// R23C55 -> BC23
		if text[0] == 'R' && unicode.IsDigit(rune(text[1])) && 1 < pos && pos < len(text) {
			row, _ := strconv.Atoi(text[1:pos])
			col, _ := strconv.Atoi(text[pos+1:])

			arrText := []rune{}
			for col > 0 {
				if col%26 == 0 {
					arrText = append(arrText, 'Z')
					col = (col - 1) / 26
				} else {
					arrText = append(arrText, rune('A'+(col%26-1)))
					col /= 26
				}
			}

			for len(arrText) > 0 {
				Fprintf("%c", arrText[len(arrText)-1])
				arrText = arrText[:len(arrText)-1]
			}

			Fprintf("%d\n", row)

		} else {
			// BC23 -> R23C55
			pos := 0
			for pos < len(text) {
				if unicode.IsDigit(rune(text[pos])) {
					break
				}
				pos++
			}

			row := text[:pos]
			col := text[pos:]

			c := 0
			for _, r := range row {
				c = c*26 + int(r-'A'+1)
			}

			Fprintf("R%sC%d\n", col, c)
		}
	}
}

func main() {
	Solution()
	writer.Flush()
}
```
