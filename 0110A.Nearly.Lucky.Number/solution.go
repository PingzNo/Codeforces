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
	var number string
	Fscanf("%s\n", &number)

	luckyCount := 0
	for _, r := range number {
		if r == '4' || r == '7' {
			luckyCount++
		}
	}

	isNearlyLucky := luckyCount > 0

	for luckyCount > 0 {
		remainder := luckyCount % 10
		if remainder != 4 && remainder != 7 {
			isNearlyLucky = false
			break
		}

		luckyCount /= 10
	}

	if isNearlyLucky {
		Fprintf("YES\n")
	} else {
		Fprintf("NO\n")
	}
}

func main() {
	Solution()
	writer.Flush()
}
