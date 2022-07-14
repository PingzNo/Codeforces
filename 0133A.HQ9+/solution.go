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
	var line string
	Fscanf("%s\n", &line)

	for _, c := range line {
		if c == 'H' || c == 'Q' || c == '9' {
			Fprintf("YES\n")
			return
		}
	}

	Fprintf("NO\n")
}

func main() {
	Solution()
	writer.Flush()
}
