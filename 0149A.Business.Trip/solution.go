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
	var centimeters int
	Fscanf("%d\n", &centimeters)

	nums := make([]int, 12)
	for numIndex := 0; numIndex < len(nums); numIndex++ {
		Fscanf("%d", &nums[numIndex])
	}

	sort.Ints(nums)

	var monthCount int
	for numIndex := len(nums) - 1; numIndex >= 0; numIndex-- {
		if centimeters <= 0 {
			break
		}

		centimeters -= nums[numIndex]
		monthCount++
	}

	if centimeters > 0 {
		Fprintf("-1\n")
	} else {
		Fprintf("%d\n", monthCount)
	}
}

func main() {
	Solution()
	writer.Flush()
}
