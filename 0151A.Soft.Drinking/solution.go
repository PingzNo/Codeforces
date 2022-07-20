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
	var n, k, l, c, d, p, nl, np int
	Fscanf("%d %d %d %d %d %d %d %d\n", &n, &k, &l, &c, &d, &p, &nl, &np)

	toastCount := k * l / nl
	sliceCount := c * d
	saltCount := p / np

	result := toastCount
	if sliceCount < result {
		result = sliceCount
	}
	if saltCount < result {
		result = saltCount
	}

	result /= n

	Fprintf("%d\n", result)
}

func main() {
	Solution()
	writer.Flush()
}
