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
	var n, m int
	Fscanf("%d %d\n", &n, &m)

	pairCount := 0
	for a := 0; a <= n; a++ {
		b := n - a*a
		if b < 0 {
			continue
		}

		if a+b*b == m {
			pairCount++
		}
	}

	Fprintf("%d\n", pairCount)
}

func main() {
	Solution()
	writer.Flush()
}
