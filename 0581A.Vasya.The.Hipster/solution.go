package main

import (
	"bufio"
	"fmt"
	"os"
	"math"
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
	var redCount, blueCount float64
	Fscanf("%f %f\n", &redCount, &blueCount)

	fashionCount := int32(math.Min(redCount, blueCount))
	pairCount := (int32(math.Max(redCount, blueCount)) - fashionCount) / 2
	Fprintf("%d %d\n", fashionCount, pairCount)
}

func main() {
	Solution()
	writer.Flush()
}
