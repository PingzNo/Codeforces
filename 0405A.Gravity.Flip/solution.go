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
	var colCount int
	Fscanf("%d\n", &colCount)

	rowCounters := [100]int{}

	var rowCount int
	for colIndex := 0; colIndex < colCount; colIndex++ {
		Fscanf("%d", &rowCount)
		for rowIndex := 0; rowIndex < rowCount; rowIndex++ {
			rowCounters[rowIndex]++
		}
	}

	for colIndex := 0; colIndex < colCount; colIndex++ {
		rowCount = 0
		for rowIndex := 0; rowIndex < len(rowCounters); rowIndex++ {
			if rowCounters[rowIndex] >= colCount-colIndex {
				rowCount++
			}

			if rowCounters[rowIndex] == 0 {
				break
			}
		}

		if colIndex > 0 {
			Fprintf(" ")
		}

		Fprintf("%d", rowCount)
	}

	Fprintf("\n")
}

func main() {
	Solution()
	writer.Flush()
}
