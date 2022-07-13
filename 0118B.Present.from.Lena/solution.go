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
	var n int
	Fscanf("%d\n", &n)

	for i := 0; i <= n; i++ {
		for j := n - i; j > 0; j-- {
			Fprintf("  ")
		}

		if i != 0 {
			for j := 0; j <= i; j++ {
				Fprintf("%d ", j)
			}
		}

		for j := i - 1; j > 0; j-- {
			Fprintf("%d ", j)
		}

		Fprintf("0\n")
	}

	for i := n - 1; i >= 0; i-- {
		for j := 1; j <= n-i; j++ {
			Fprintf("  ")
		}

		if i != 0 {
			for j := 0; j <= i; j++ {
				Fprintf("%d ", j)
			}
		}

		for j := i - 1; j > 0; j-- {
			Fprintf("%d ", j)
		}

		Fprintf("0\n")
	}
}

func main() {
	Solution()
	writer.Flush()
}
