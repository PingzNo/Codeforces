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
	var rowCount, colCount int
	Fscanf("%d %d\n", &rowCount, &colCount)

	toRight := true
	for rowIndex := 0; rowIndex < rowCount; rowIndex++ {
		if (rowIndex+1)%2 == 1 {
			for colIndex := 0; colIndex < colCount; colIndex++ {
				Fprintf("#")
			}
		} else {
			if toRight {
				for colIndex := 1; colIndex < colCount; colIndex++ {
					Fprintf(".")
				}
				Fprintf("#")
			} else {
				Fprintf("#")
				for colIndex := 1; colIndex < colCount; colIndex++ {
					Fprintf(".")
				}
			}
			toRight = !toRight
		}
		Fprintf("\n")
	}
}

func main() {
	Solution()
	writer.Flush()
}
