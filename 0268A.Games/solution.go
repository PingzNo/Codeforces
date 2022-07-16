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
	var teamCount int
	Fscanf("%d\n", &teamCount)

	hostCounter := map[int]int{}
	guestCounter := map[int]int{}

	var hostColor, guestColor int
	for teamIndex := 0; teamIndex < teamCount; teamIndex++ {
		Fscanf("%d %d\n", &hostColor, &guestColor)
		hostCounter[hostColor]++
		guestCounter[guestColor]++
	}

	result := 0
	for hostColor, hostCount := range hostCounter {
		result += hostCount * guestCounter[hostColor]
	}

	Fprintf("%d\n", result)
}

func main() {
	Solution()
	writer.Flush()
}
