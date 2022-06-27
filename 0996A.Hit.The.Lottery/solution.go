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
	var money int
	Fscanf("%d\n", &money)

	dividers := [5]int{100, 20, 10, 5, 1}

	billCountSum := 0
	for _, divider := range dividers {
		billCount := money / divider
		billCountSum += billCount
		money -= billCount * divider

	}

	Fprintf("%d\n", billCountSum)
}

func main() {
	Solution()
	writer.Flush()
}
