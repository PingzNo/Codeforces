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
	var nameCount int
	Fscanf("%d\n", &nameCount)

	names := map[string]int{}

	var name string
	for nameIndex := 0; nameIndex < nameCount; nameIndex++ {
		Fscanf("%s\n", &name)
		if count, nameFound := names[name]; !nameFound {
			names[name] = 1
			Fprintf("OK\n")
		} else {
			names[name]++
			Fprintf("%s%d\n", name, count)
		}
	}
}

func main() {
	Solution()
	writer.Flush()
}
