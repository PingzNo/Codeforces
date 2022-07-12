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
	var bottleCount int
	Fscanf("%d\n", &bottleCount)

	var percent, percentSum int
	for bottleIndex := 0; bottleIndex < bottleCount; bottleIndex++ {
		Fscanf("%d", &percent)
		percentSum += percent
	}

	Fprintf("%.5f\n", float64(percentSum)/float64(bottleCount))
}

func main() {
	Solution()
	writer.Flush()
}
