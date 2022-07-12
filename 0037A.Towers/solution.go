package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	var barCount int
	Fscanf("%d\n", &barCount)

	barHeights := make([]int, barCount)
	for barIndex := 0; barIndex < barCount; barIndex++ {
		Fscanf("%d", &barHeights[barIndex])
	}

	sort.Ints(barHeights)

	height := 1
	maxLength := -1

	var i, j int
	for i = 0; i < barCount; i++ {
		for j = i + 1; j < barCount; j++ {
			if barHeights[i] != barHeights[j] {
				height++
				break
			}
		}

		length := j - i
		if length > maxLength {
			maxLength = length
		}

		i = j - 1
	}

	Fprintf("%d %d\n", maxLength, height)
}

func main() {
	Solution()
	writer.Flush()
}
