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
	var soldierCount int
	Fscanf("%d\n", &soldierCount)

	minHeight := 101
	minIndex := -1
	maxHeight := 0
	maxIndex := -1

	var height int
	for soldierIndex := 0; soldierIndex < soldierCount; soldierIndex++ {
		Fscanf("%d", &height)
		if height <= minHeight {
			minHeight = height
			minIndex = soldierIndex
		}

		if height > maxHeight {
			maxHeight = height
			maxIndex = soldierIndex
		}
	}

	results := (soldierCount - minIndex) + (maxIndex - 1)
	if minIndex < maxIndex {
		results--
	}

	Fprintf("%d\n", results)
}

func main() {
	Solution()
	writer.Flush()
}
