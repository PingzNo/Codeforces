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
	var pos int
	Fscanf("%d\n", &pos)

	stepCount := pos / 5
	if pos%5 != 0 {
		stepCount++
	}

	Fprintf("%d\n", stepCount)
}

func main() {
	Solution()
	writer.Flush()
}
