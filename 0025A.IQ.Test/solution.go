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
	var numberCount int
	Fscanf("%d\n", &numberCount)

	counters := [2]int{}
	indices := [2]int{}

	var number int
	for numberIndex := 0; numberIndex < numberCount; numberIndex++ {
		Fscanf("%d", &number)
		index := number % 2
		counters[index]++
		if counters[index] == 1 {
			indices[index] = numberIndex + 1
		}

		if counters[index] > 1 && counters[1-index] == 1 {
			Fprintf("%d\n", indices[1-index])
			break
		}

		if counters[index] == 1 && counters[1-index] > 1 {
			Fprintf("%d\n", indices[index])
			break
		}
	}
}

func main() {
	Solution()
	writer.Flush()
}
