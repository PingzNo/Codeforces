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
	var song string
	Fscanf("%s\n", &song)

	isNewWord := true
	isSongStarted := false

	i := 0
	for i < len(song) {
		if i < len(song)-2 && song[i] == 'W' && song[i+1] == 'U' && song[i+2] == 'B' {
			i += 3
			isNewWord = true
		} else {
			if isNewWord {
				if isSongStarted {
					Fprintf(" ")
				} else {
					isSongStarted = true
				}

				isNewWord = false
			}

			Fprintf("%c", song[i])
			i++
		}
	}
	Fprintf("\n")
}

func main() {
	Solution()
	writer.Flush()
}
