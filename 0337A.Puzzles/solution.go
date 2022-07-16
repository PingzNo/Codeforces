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
	var childCount, puzzleCount int
	Fscanf("%d %d\n", &childCount, &puzzleCount)

	puzzles := make([]int, puzzleCount)
	for i := 0; i < puzzleCount; i++ {
		Fscanf("%d", &puzzles[i])
	}

	sort.Ints(puzzles)

	minDiff := 1000
	for i := 0; i+childCount-1 < puzzleCount; i++ {
		if puzzles[i+childCount-1]-puzzles[i] < minDiff {
			minDiff = puzzles[i+childCount-1] - puzzles[i]
		}
	}

	Fprintf("%d\n", minDiff)
}

func main() {
	Solution()
	writer.Flush()
}
