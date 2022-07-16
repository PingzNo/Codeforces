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
	var year int
	Fscanf("%d\n", &year)

	var a, b, c, d int
	for a = year / 1000; a <= 9; a++ {
		for b = 0; b <= 9; b++ {
			if b == a {
				continue
			}

			for c = 0; c <= 9; c++ {
				if c == a || c == b {
					continue
				}

				for d = 0; d <= 9; d++ {
					if d == a || d == b || d == c {
						continue
					}

					prefectYear := a*1000 + b*100 + c*10 + d
					if prefectYear > year {
						Fprintf("%d\n", prefectYear)
						return
					}
				}
			}
		}
	}
}

func main() {
	Solution()
	writer.Flush()
}
