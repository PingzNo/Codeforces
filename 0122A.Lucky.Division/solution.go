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
	var number int
	Fscanf("%d\n", &number)

	dividers := []int{4, 7, 47, 74, 447, 474, 477, 744, 747, 774}
	for _, divider := range dividers {
		if divider <= number {
			if number%divider == 0 {
				Fprintf("YES\n")
				return
			}
		} else {
			break
		}
	}

	Fprintf("NO\n")
}

func main() {
	Solution()
	writer.Flush()
}
