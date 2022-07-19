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
	var numCount int
	Fscanf("%d\n", &numCount)

	var num int64
	scores := make([]int64, 1e5+1)
	for numIndex := 0; numIndex < numCount; numIndex++ {
		Fscanf("%d", &num)
		scores[num] += num
	}

	for numIndex := 2; numIndex < len(scores); numIndex++ {
		if scores[numIndex-1] < scores[numIndex-2]+scores[numIndex] {
			scores[numIndex] += scores[numIndex-2]
		} else {
			scores[numIndex] = scores[numIndex-1]
		}
	}

	Fprintf("%d\n", scores[len(scores)-1])
}

func main() {
	Solution()
	writer.Flush()
}
