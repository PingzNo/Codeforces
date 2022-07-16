package main

import (
	"bufio"
	"fmt"
	"math"
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
	var number int
	for rowIndex := 0; rowIndex < 5; rowIndex++ {
		for colIndex := 0; colIndex < 5; colIndex++ {
			Fscanf("%d", &number)
			if number == 1 {
				rowDistance := int(math.Abs(float64(rowIndex - 2)))
				colDistance := int(math.Abs(float64(colIndex - 2)))
				Fprintf("%d\n", rowDistance+colDistance)
				return
			}
		}
		Fscanf("\n")
	}
}

func main() {
	Solution()
	writer.Flush()
}
