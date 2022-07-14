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
	var friendCount int
	Fscanf("%d\n", &friendCount)

	giftFroms := make([]int, friendCount)

	var giftTo int
	for friendIndex := 0; friendIndex < friendCount; friendIndex++ {
		Fscanf("%d", &giftTo)
		giftFroms[giftTo-1] = friendIndex + 1
	}

	for friendIndex, giftFrom := range giftFroms {
		if friendIndex > 0 {
			Fprintf(" %d", giftFrom)
		} else {
			Fprintf("%d", giftFrom)
		}
	}
	Fprintf("\n")
}

func main() {
	Solution()
	writer.Flush()
}
