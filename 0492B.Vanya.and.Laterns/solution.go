package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	var latternCount, length int
	Fscanf("%d %d\n", &latternCount, &length)

	pos := make([]float64, latternCount)
	for i := 0; i < latternCount; i++ {
		Fscanf("%f", &pos[i])
	}

	sort.Float64s(pos)

	var minRadius float64
	for i := 0; i < latternCount; i++ {
		if pos[i] == 0 {
			continue
		}

		if i == 0 {
			minRadius = pos[i]
			continue
		}

		newRadius := (pos[i] - pos[i-1]) / 2
		if newRadius > minRadius {
			minRadius = newRadius
		}
	}

	if pos[latternCount-1] != float64(length) {
		newRadius := float64(length) - pos[latternCount-1]
		if newRadius > minRadius {
			minRadius = newRadius
		}
	}

	Fprintf("%.10f\n", minRadius)
}

func main() {
	Solution()
	writer.Flush()
}
