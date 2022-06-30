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
	var friendCount, fenceHeight int
	Fscanf("%d %d\n", &friendCount, &fenceHeight)

	var friendHeight, roadWidth int
	for friendIndex := 0; friendIndex < friendCount; friendIndex++ {
		Fscanf("%d", &friendHeight)
		if friendHeight > fenceHeight {
			roadWidth += 2
		} else {
			roadWidth++
		}
	}

	Fprintf("%d\n", roadWidth)
}

func main() {
	Solution()
	writer.Flush()
}
