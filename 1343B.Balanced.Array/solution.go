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
	var caseCount int
	Fscanf("%d\n", &caseCount)

	var numCount int
	for caseIndex := 0; caseIndex < caseCount; caseIndex++ {
		Fscanf("%d\n", &numCount)

		halfCount := numCount / 2
		if halfCount%2 == 1 {
			Fprintf("NO\n")
			continue
		}

		halfHalfCount := halfCount / 2
		refNum := 0

		var result []int
		for numIndex := halfHalfCount; numIndex > 0; numIndex-- {
			result = append(result, refNum-numIndex*2)
		}

		for numIndex := 0; numIndex < halfHalfCount; numIndex++ {
			result = append(result, refNum+numIndex*2)
		}

		for numIndex := halfHalfCount; numIndex > 0; numIndex-- {
			result = append(result, refNum-numIndex*2-1)
		}

		for numIndex := 0; numIndex < halfHalfCount; numIndex++ {
			result = append(result, refNum+numIndex*2+1)
		}

		first := result[0]
		for numIndex := 0; numIndex < len(result); numIndex++ {
			result[numIndex] = result[numIndex] - first + 2
		}

		Fprintf("YES\n")
		for i, x := range result {
			if i > 0 {
				Fprintf(" ")
			}

			Fprintf("%d", x)
		}
		Fprintf("\n")
	}
}

func main() {
	Solution()
	writer.Flush()
}
