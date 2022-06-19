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
	var houseCount, taskCount int64
	Fscanf("%d %d\n", &houseCount, &taskCount)

	currentIndex := int64(1)

	var houseIndex, timeCost int64
	for taskIndex := int64(0); taskIndex < taskCount; taskIndex++ {
		Fscanf("%d", &houseIndex)
		if currentIndex <= houseIndex {
			timeCost += houseIndex - currentIndex
		} else {
			timeCost += houseCount - currentIndex + houseIndex
		}

		currentIndex = houseIndex
	}

	Fprintf("%d\n", timeCost)
}

func main() {
	Solution()
	writer.Flush()
}
