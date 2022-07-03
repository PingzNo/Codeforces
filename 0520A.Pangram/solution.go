package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	var letterCount int
	var letters string
	Fscanf("%d\n%s\n", &letterCount, &letters)

	letterMap := make(map[rune]bool)
	for _, letter := range strings.ToUpper(letters) {
		letterMap[letter] = true
	}

	if len(letterMap) == 26 {
		Fprintf("YES\n")
		return
	}

	Fprintf("NO\n")
}

func main() {
	Solution()
	writer.Flush()
}
