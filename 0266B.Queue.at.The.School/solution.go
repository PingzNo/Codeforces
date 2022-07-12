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
	var headCount, shiftCount int
	Fscanf("%d %d\n", &headCount, &shiftCount)

	var line string
	Fscanf("%s\n", &line)

	firstQueue := []rune(line)
	secondQueue := []rune(line)

	p := &firstQueue
	q := &secondQueue

	for shiftIndex := 0; shiftIndex < shiftCount; shiftIndex++ {
		queueIndex := 1
		for queueIndex < headCount {
			if (*p)[queueIndex-1] == 'B' && (*p)[queueIndex] == 'G' {
				(*q)[queueIndex-1], (*q)[queueIndex] = 'G', 'B'
				if queueIndex < headCount-1 && queueIndex+2 >= headCount {
					(*q)[headCount-1] = (*p)[headCount-1]
				}
				queueIndex += 2
			} else {
				(*q)[queueIndex-1], (*q)[queueIndex] = (*p)[queueIndex-1], (*p)[queueIndex]
				queueIndex++
			}
		}

		p, q = q, p
	}

	for _, r := range *p {
		Fprintf("%c", r)
	}
	Fprintf("\n")
}

func main() {
	Solution()
	writer.Flush()
}
