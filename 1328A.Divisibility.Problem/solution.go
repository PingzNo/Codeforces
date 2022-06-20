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
	var caseCount int
	Fscanf("%d\n", &caseCount)

	var a, b int
	for caseIndex := 0; caseIndex < caseCount; caseIndex++ {
		Fscanf("%d %d\n", &a, &b)
		if a%b == 0 {
			Fprintf("0\n")
		} else {
			Fprintf("%d\n", b-a%b)
		}
	}
}

func main() {
	Solution()
	writer.Flush()
}
