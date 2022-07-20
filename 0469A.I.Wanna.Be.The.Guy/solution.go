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
	var levelCount int
	Fscanf("%d\n", &levelCount)

	levelPassed := map[int]bool{}

	var playCount, level int
	for personIndex := 0; personIndex < 2; personIndex++ {
		Fscanf("%d", &playCount)
		for playIndex := 0; playIndex < playCount; playIndex++ {
			Fscanf("%d", &level)
			levelPassed[level-1] = true
		}

		Fscanf("\n")
	}

	if len(levelPassed) == levelCount {
		Fprintf("I become the guy.\n")
	} else {
		Fprintf("Oh, my keyboard!\n")
	}

}

func main() {
	Solution()
	writer.Flush()
}
