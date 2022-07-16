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
	var contestCount int
	Fscanf("%d\n", &contestCount)

	minScore := -1
	maxScore := -1
	awesomeCount := 0

	var score int
	for contestIndex := 0; contestIndex < contestCount; contestIndex++ {
		Fscanf("%d", &score)
		if minScore == -1 {
			minScore = score
		} else if score < minScore {
			minScore = score
			awesomeCount++
		}

		if maxScore == -1 {
			maxScore = score
		} else if maxScore < score {
			maxScore = score
			awesomeCount++
		}
	}

	Fprintf("%d\n", awesomeCount)
}

func main() {
	Solution()
	writer.Flush()
}
