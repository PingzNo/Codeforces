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
	var n, m, a, b int
	Fscanf("%d %d %d %d\n", &n, &m, &a, &b)

	aCost := n * a
	bCost := n/m*b + int(math.Min(float64(n%m*a), float64(b)))
	Fprintf("%d\n", int(math.Min(float64(aCost), float64(bCost))))
}

func main() {
	Solution()
	writer.Flush()
}
