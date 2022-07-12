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
	var number int64
	Fscanf("%d\n", &number)

	var result int64
	if number % 2 == 0 {
		result = number/2
	} else {
		result = number/2 - number
	}

	Fprintf("%d\n", result)
}

func main() {
	Solution()
	writer.Flush()
}
