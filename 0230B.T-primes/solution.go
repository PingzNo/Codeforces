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
	count := int64(10e6)
	flags := make([]bool, count+1)
	flags[0] = true
	flags[1] = true

	upperLimit := int64(math.Sqrt(float64(count))) + 1
	for divider := int64(2); divider <= upperLimit; divider++ {
		if flags[divider] {
			continue
		}

		for x := divider + divider; x <= int64(count); x += divider {
			flags[x] = true
		}
	}

	var numberCount int
	Fscanf("%d\n", &numberCount)

	var number int64
	for numberIndex := 0; numberIndex < numberCount; numberIndex++ {
		Fscanf("%d", &number)
		root := int64(math.Sqrt(float64(number)))
		if root*root == number && !flags[root] {
			Fprintf("YES\n")
			continue
		}

		Fprintf("NO\n")
	}
}

func main() {
	Solution()
	writer.Flush()
}
