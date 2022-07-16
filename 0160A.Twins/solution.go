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
	var coinCount int
	Fscanf("%d\n", &coinCount)

	coinValues := make([]int, coinCount)

	coinSum := 0
	for coinIndex := 0; coinIndex < coinCount; coinIndex++ {
		Fscanf("%d", &coinValues[coinIndex])
		coinSum += coinValues[coinIndex]
	}

	sort.Ints(coinValues)

	leftSum := coinSum
	coinIndex := coinCount - 1
	for ; coinIndex >= 0; coinIndex-- {
		leftSum -= coinValues[coinIndex]
		if leftSum < coinSum-leftSum {
			break
		}
	}

	Fprintf("%d\n", coinCount-coinIndex)
}

func main() {
	Solution()
	writer.Flush()
}
