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
	var magCount int
	Fscanf("%d\n", &magCount)

	var first byte = 'X'
	groupCount := 0
	var mag string
	for magIndex := 0; magIndex < magCount; magIndex++ {
		Fscanf("%s\n", &mag)
		if first != mag[0] {
			groupCount++
			first = mag[0]
		}
	}

	Fprintf("%d\n", groupCount)
}

func main() {
	Solution()
	writer.Flush()
}
