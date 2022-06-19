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
	var n, m int
	Fscanf("%d %d\n", &n, &m)

	var minValue int
	if n < m {
		minValue = n
	} else {
		minValue = m
	}

	if minValue%2 == 1 {
		Fprintf("Akshat\n")
	} else {
		Fprintf("Malvika\n")
	}
}

func main() {
	Solution()
	writer.Flush()
}
