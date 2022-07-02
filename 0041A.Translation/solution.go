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
	var berlandish, birlandish string
	Fscanf("%s\n%s\n", &berlandish, &birlandish)

	if len(berlandish) != len(birlandish) {
		Fprintf("NO\n")
		return
	}

	for i := 0; i < len(berlandish); i++ {
		if berlandish[i] != birlandish[len(birlandish)-i-1] {
			Fprintf("NO\n")
			return
		}
	}

	Fprintf("YES\n")
}

func main() {
	Solution()
	writer.Flush()
}
