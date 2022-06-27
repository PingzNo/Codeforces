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
	var rooCount int
	Fscanf("%d\n", &rooCount)

	var roomToMoveIn int
	var peopleCount, roomCap int
	for roomIndex := 0; roomIndex < rooCount; roomIndex++ {
		Fscanf("%d %d\n", &peopleCount, &roomCap)
		if roomCap-peopleCount >= 2 {
			roomToMoveIn++
		}
	}

	Fprintf("%d\n", roomToMoveIn)
}

func main() {
	Solution()
	writer.Flush()
}
