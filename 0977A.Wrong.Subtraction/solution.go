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
	var number, times int
	Fscanf("%d %d\n", &number, &times)

	remainder := number % 10
	for number > 0 && times > 0 {
		if remainder == 0 {
			times--
			number /= 10

		} else if remainder <= times {
			times -= remainder
			number -= remainder
		} else {
			number -= times
			times = 0
		}

		remainder = number % 10
	}

	Fprintf("%d\n", number)
}

func main() {
	Solution()
	writer.Flush()
}
