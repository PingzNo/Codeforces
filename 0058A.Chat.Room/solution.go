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
