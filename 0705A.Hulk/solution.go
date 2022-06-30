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
	var layerCount int
	Fscanf("%d\n", &layerCount)

	actions := []string{"hate", "love"}
	actionIndex := 0

	for layerIndex := 1; layerIndex <= layerCount; layerIndex++ {
		if layerIndex == layerCount {
			Fprintf("I %s it\n", actions[actionIndex])
		} else {
			Fprintf("I %s that ", actions[actionIndex])
		}

		actionIndex = (actionIndex + 1) % 2
	}
}

func main() {
	Solution()
	writer.Flush()
}
