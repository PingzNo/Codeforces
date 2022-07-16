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
	var word string
	Fscanf("%s\n", &word)
	Fprintf("%c%s\n", unicode.ToUpper(rune(word[0])), word[1:])
}

func main() {
	Solution()
	writer.Flush()
}
