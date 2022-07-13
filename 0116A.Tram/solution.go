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
	var stopCount int
	Fscanf("%d\n", &stopCount)

	headCount := 0
	maxCapacity := 0

	var outCount, inCount int
	for stopIndex := 0; stopIndex < stopCount; stopIndex++ {
		Fscanf("%d %d\n", &outCount, &inCount)
		headCount += inCount - outCount
		if headCount > maxCapacity {
			maxCapacity = headCount
		}
	}

	Fprintf("%d\n", maxCapacity)
}

func main() {
	Solution()
	writer.Flush()
}
