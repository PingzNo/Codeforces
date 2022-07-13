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
	var text string
	Scanf("%s\n", &text)
	text = strings.ToLower(text)

	for i := 0; i < len(text); i++ {
		switch text[i] {
		case 'a':
		case 'o':
		case 'y':
		case 'e':
		case 'u':
		case 'i':
			continue

		default:
			Printf(".%c", text[i])
			break
		}
	}

	Printf("\n")
}

func main() {
	Solution()
	writer.Flush()
}
