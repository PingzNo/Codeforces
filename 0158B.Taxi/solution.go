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
	var group, groupCount int
	Fscanf("%d\n", &groupCount)

	counter := [4]int{}
	for groupIndex := 0; groupIndex < groupCount; groupIndex++ {
		Fscanf("%d", &group)
		counter[group-1]++
	}

	taxiCount := counter[3]

	taxiCount += counter[2]
	if counter[2] <= counter[0] {
		counter[0] -= counter[2]
	} else {
		counter[0] = 0
	}

	taxiCount += counter[1] / 2
	counter[1] %= 2
	if counter[1] == 1 {
		taxiCount++
		if counter[0] >= 2 {
			counter[0] -= 2
		} else if counter[0] >= 1 {
			counter[0] -= 1
		}
	}

	taxiCount += counter[0] / 4
	if counter[0]%4 > 0 {
		taxiCount++
	}

	Fprintf("%d\n", taxiCount)
}

func main() {
	Solution()
	writer.Flush()
}
