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
	var stepCount, prevIncome, income, beginIndex int
	Fscanf("%d\n%d", &stepCount, &prevIncome)

	maxLength := 1
	stepIndex := 1
	for stepIndex < stepCount {
		Fscanf("%d", &income)
		if income < prevIncome {
			if stepIndex-beginIndex+1 > maxLength {
				maxLength = stepIndex - beginIndex
			}

			beginIndex = stepIndex
		}

		prevIncome = income
		stepIndex++
	}

	if stepIndex-beginIndex+1 > maxLength {
		maxLength = stepIndex - beginIndex
	}

	Fprintf("%d\n", maxLength)
}

func main() {
	Solution()
	writer.Flush()
}
