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
