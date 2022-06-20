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
