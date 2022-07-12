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
	var x, y string
	Fscanf("%s\n%s\n", &x, &y)

	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			Fprintf("1")
		} else {
			Fprintf("0")
		}
	}
	Fprintf("\n")
}

func main() {
	Solution()
	writer.Flush()
}
