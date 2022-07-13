package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Printf(format string, x ...interface{}) {
	fmt.Fprintf(writer, format, x...)
}

func Scanf(format string, x ...interface{}) {
	fmt.Fscanf(reader, format, x...)
}

func Solution() {
	var first, second string
	Scanf("%s\n%s\n", &first, &second)
	first = strings.ToUpper(first)
	second = strings.ToUpper(second)

	result := 0
	for i := 0; i < len(first); i++ {
		if first[i] < second[i] {
			result = -1
			break
		}

		if first[i] > second[i] {
			result = 1
			break
		}
	}

	Printf("%d\n", result)
}

func main() {
	Solution()
	writer.Flush()
}
