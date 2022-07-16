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
	var name string
	Fscanf("%s\n", &name)

	haveRunes := [26]int{}
	for _, r := range name {
		haveRunes[r-'a'] = 1
	}

	sum := 0
	for _, haveRune := range haveRunes {
		sum += haveRune
	}

	if sum%2 == 0 {
		Fprintf("CHAT WITH HER!\n")
	} else {
		Fprintf("IGNORE HIM!\n")
	}
}

func main() {
	Solution()
	writer.Flush()
}
