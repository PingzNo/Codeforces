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
	var guest, host, pile string
	Fscanf("%s\n%s\n%s\n", &guest, &host, &pile)

	counter := [26]int{}
	for _, r := range guest {
		counter[r-'A']--
	}

	for _, r := range host {
		counter[r-'A']--
	}

	for _, r := range pile {
		counter[r-'A']++
	}

	for _, count := range counter {
		if count != 0 {
			Fprintf("NO\n")
			return
		}
	}

	Fprintf("YES\n")
}

func main() {
	Solution()
	writer.Flush()
}
