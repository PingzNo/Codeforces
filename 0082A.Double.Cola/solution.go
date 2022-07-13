package main

import (
	"bufio"
	"fmt"
	"math"
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
	names := []string{
		"Sheldon",
		"Leonard",
		"Penny",
		"Rajesh",
		"Howard",
	}

	var rank int
	Fscanf("%d\n", &rank)

	accumulation := 0
	nameCount := len(names)
	prevAccumulation := 0
	for power := 1; ; power++ {
		accumulation += nameCount
		if rank > accumulation {
			prevAccumulation = accumulation
			nameCount += nameCount
			continue
		}

		nameIndex := (rank - prevAccumulation - 1) / int(math.Pow(2.0, float64(power-1))+0.5)
		Fprintf("%s\n", names[nameIndex])
		break
	}
}

func main() {
	Solution()
	writer.Flush()
}
