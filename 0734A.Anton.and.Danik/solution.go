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
	var gameCount int
	var gameRecords string
	Fscanf("%d\n%s\n", &gameCount, &gameRecords)

	var antonCount, danikCount int
	for _, record := range gameRecords {
		if record == 'A' {
			antonCount++
		} else if record == 'D' {
			danikCount++
		}
	}

	if antonCount > danikCount {
		Fprintf("Anton\n")
	} else if danikCount > antonCount {
		Fprintf("Danik\n")
	} else {
		Fprintf("Friendship\n")
	}
}

func main() {
	Solution()
	writer.Flush()
}
