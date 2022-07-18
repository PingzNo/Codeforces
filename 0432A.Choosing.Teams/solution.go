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
	var personCount, minContestCount int
	Fscanf("%d %d\n", &personCount, &minContestCount)

	var contestCount, candidateCount int
	for personIndex := 0; personIndex < personCount; personIndex++ {
		Fscanf("%d", &contestCount)
		if 5-contestCount >= minContestCount {
			candidateCount++
		}
	}
	Fscanf("\n")

	Fprintf("%d\n", candidateCount/3)
}

func main() {
	Solution()
	writer.Flush()
}
