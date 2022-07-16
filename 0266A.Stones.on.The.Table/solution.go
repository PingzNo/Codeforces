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
	var stoneCount int
	var stones string
	Fscanf("%d\n%s\n", &stoneCount, &stones)

	removedCount := 0
	stone := stones[0]
	for stoneIndex := 1; stoneIndex < stoneCount; stoneIndex++ {
		if stones[stoneIndex] == stone {
			removedCount++
		} else {
			stone = stones[stoneIndex]
		}
	}

	Fprintf("%d\n", removedCount)
}

func main() {
	Solution()
	writer.Flush()
}
