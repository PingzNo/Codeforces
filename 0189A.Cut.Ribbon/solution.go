package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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
	var ribbonLen int
	cutLen := make([]int, 3)
	Fscanf("%d %d %d %d\n", &ribbonLen, &cutLen[0], &cutLen[1], &cutLen[2])

	sort.Ints(cutLen)

	cutNumber := len(cutLen)
	if cutLen[2]%cutLen[1] == 0 || cutLen[2]%cutLen[0] == 0 || cutLen[2]%(cutLen[0]+cutLen[1]) == 0 || cutLen[2] > ribbonLen {
		cutNumber = 2
	}

	if cutLen[1]%cutLen[0] == 0 || cutLen[1] > ribbonLen {
		if cutNumber == 3 {
			cutLen[1] = cutLen[2]
			cutNumber = 2
		} else {
			Fprintf("%d\n", ribbonLen/cutLen[0])
			return
		}
	}

	table := make([]int, ribbonLen+1)
	for cutIndex := 0; cutIndex < cutNumber; cutIndex++ {
		table[cutLen[cutIndex]] = 1
	}

	for cutIndex := 0; cutIndex < cutNumber; cutIndex++ {
		for j := cutLen[cutIndex]; j <= ribbonLen; j++ {
			if table[j-cutLen[cutIndex]] > 0 {
				table[j] = int(math.Max(float64(table[j-cutLen[cutIndex]]+1), float64(table[j])))
			}
		}
	}

	Fprintf("%d\n", table[ribbonLen])
}

func main() {
	Solution()
	writer.Flush()
}
