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
	var peopleCount, a, b int
	Fscanf("%d %d %d\n", &peopleCount, &a, &b)

	lowerBound := a + 1
	upperBound := peopleCount
	if lowerBound > upperBound {
		Fprintf("0\n")
		return
	}

	if upperBound-lowerBound > b {
		lowerBound = peopleCount - b
	}

	Fprintf("%d\n", upperBound-lowerBound+1)
}

func main() {
	Solution()
	writer.Flush()
}
