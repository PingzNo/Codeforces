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
	var cubeCount int
	Fscanf("%d\n", &cubeCount)

	faceCount := 0

	var cubeName string
	for cubeIndex := 0; cubeIndex < cubeCount; cubeIndex++ {
		Fscanf("%s\n", &cubeName)
		switch cubeName[0] {
		case 'T':
			faceCount += 4
		case 'C':
			faceCount += 6
		case 'O':
			faceCount += 8
		case 'D':
			faceCount += 12
		case 'I':
			faceCount += 20
		}
	}

	Fprintf("%d\n", faceCount)
}

func main() {
	Solution()
	writer.Flush()
}
