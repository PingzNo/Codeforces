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
	var a, b, c int
	Fscanf("%d\n%d\n%d\n", &a, &b, &c)

	var maxValue, result int

	result = a + b + c
	if result > maxValue {
		maxValue = result
	}

	result = a * b * c
	if result > maxValue {
		maxValue = result
	}

	result = (a + b) * c
	if result > maxValue {
		maxValue = result
	}

	result = a * (b + c)
	if result > maxValue {
		maxValue = result
	}

	result = a*b + c
	if result > maxValue {
		maxValue = result
	}

	result = a + b*c
	if result > maxValue {
		maxValue = result
	}

	Fprintf("%d\n", maxValue)
}

func main() {
	Solution()
	writer.Flush()
}
