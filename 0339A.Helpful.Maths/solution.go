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

	var counts [3]int
	for _, character := range line {
		if character != '+' {
			counts[character-'1']++
		}
	}

	isFirst := true
	for index := 0; index < 3; index++ {
		for count := 0; count < counts[index]; count++ {
			if isFirst {
				isFirst = false
			} else {
				Fprintf("+")
			}
			Fprintf("%d", index+1)
		}
	}
	Fprintf("\n")
}

func main() {
	Solution()
	writer.Flush()
}
