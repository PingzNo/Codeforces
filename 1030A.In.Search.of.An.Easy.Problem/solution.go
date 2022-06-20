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
	var questionCount int
	Fscanf("%d\n", &questionCount)

	var opinion int
	for questionIndex := 0; questionIndex < questionCount; questionIndex++ {
		Fscanf("%d", &opinion)
		if opinion == 1 {
			Fprintf("HARD\n")
			return
		}
	}

	Fprintf("EASY\n")
}

func main() {
	Solution()
	writer.Flush()
}
