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
	haveShoes := map[int]bool{}

	var shoe int
	for i := 0; i < 4; i++ {
		Fscanf("%d", &shoe)
		haveShoes[shoe] = true
	}

	Fprintf("%d\n", 4-len(haveShoes))
}

func main() {
	Solution()
	writer.Flush()
}
