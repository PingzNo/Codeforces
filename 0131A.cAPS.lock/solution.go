package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	var text string
	Fscanf("%s\n", &text)

	if text == strings.ToUpper(text) || (unicode.IsLower(rune(text[0])) && text[1:] == strings.ToUpper(text[1:])) {
		for _, c := range text {
			if unicode.IsLower(c) {
				Fprintf("%c", unicode.ToUpper(c))
			} else {
				Fprintf("%c", unicode.ToLower(c))
			}
		}
		Fprintf("\n")

	} else {
		Fprintf("%s\n", text)
	}
}

func main() {
	Solution()
	writer.Flush()
}
