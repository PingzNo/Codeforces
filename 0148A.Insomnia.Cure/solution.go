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
	var divider, dragonCount int
	dividers := []int{}
	for i := 0; i < 4; i++ {
		Fscanf("%d\n", &divider)
		if divider == 1 {
			for j := i + 1; j < 5; j++ {
				Fscanf("%d\n", &dragonCount)
			}

			Fprintf("%d\n", dragonCount)
			return
		}

		isDivided := false
		for j, x := range dividers {
			if x%divider == 0 {
				dividers[j] = divider
				isDivided = true
				continue
			}

			if divider%x == 0 {
				isDivided = true
				continue
			}
		}

		if !isDivided {
			dividers = append(dividers, divider)
		}
	}

	Fscanf("%d\n", &dragonCount)

	damagedCount := 0
	for x := 1; x <= dragonCount; x++ {
		for _, y := range dividers {
			if x%y == 0 {
				damagedCount++
				break
			}
		}
	}

	Fprintf("%d\n", damagedCount)
}

func main() {
	Solution()
	writer.Flush()
}
