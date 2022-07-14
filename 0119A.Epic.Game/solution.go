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

func Gcd(x, y int) int {
	for y != 0 {
		t := y
		y = x % y
		x = t
	}

	return x
}

func Solution() {
	var stoneCount int
	simon := [2]int{}
	Fscanf("%d %d %d\n", &simon[0], &simon[1], &stoneCount)

	pointerIndex := 0
	winnerIndex := 0

	var gcdNumber int
	for {
		gcdNumber = Gcd(simon[pointerIndex], stoneCount)
		if gcdNumber > stoneCount {
			winnerIndex = 1 - pointerIndex
			break
		}

		stoneCount -= gcdNumber
		pointerIndex = 1 - pointerIndex

		if stoneCount < 0 {
			break
		}
	}

	Fprintf("%d\n", winnerIndex)
}

func main() {
	Solution()
	writer.Flush()
}
