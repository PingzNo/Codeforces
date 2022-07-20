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
	var balance int
	Fscanf("%d\n", &balance)

	if balance >= 0 {
		Fprintf("%d\n", balance)
	} else {
		balance = -balance
		if balance < 10 {
			Fprintf("0\n")
		} else {
			lastDigit := balance % 10
			digits := (balance - lastDigit) / 10
			prevDigit := digits % 10

			var result int
			if lastDigit < prevDigit {
				result = -((digits - prevDigit) + lastDigit)
			} else {
				result = -((digits - prevDigit) + prevDigit)
			}

			Fprintf("%d\n", result)
		}
	}
}

func main() {
	Solution()
	writer.Flush()
}
