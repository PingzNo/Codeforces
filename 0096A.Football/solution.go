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
	var status string
	Fscanf("%s\n", &status)

	team := '0'
	count := 0
	for _, r := range status {
		if r == team {
			count++
		} else {
			team = r
			count = 1
		}

		if count >= 7 {
			break
		}
	}

	if count >= 7 {
		Fprintf("YES\n")
	} else {
		Fprintf("NO\n")
	}
}

func main() {
	Solution()
	writer.Flush()
}
