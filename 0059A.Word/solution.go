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
