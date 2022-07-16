package main

import (
	"bufio"
	"fmt"
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

type Dragon struct {
	Strength int
	Bonus    int
}

func Solution() {
	var strength, dragonCount int
	Fscanf("%d %d\n", &strength, &dragonCount)

	dragons := make([]Dragon, dragonCount)
	for dragonIndex := 0; dragonIndex < dragonCount; dragonIndex++ {
		Fscanf("%d %d\n", &dragons[dragonIndex].Strength, &dragons[dragonIndex].Bonus)
	}

	sort.Slice(dragons, func(i, j int) bool {
		return dragons[i].Strength < dragons[j].Strength
	})

	gamePassed := true
	for dragonIndex := 0; dragonIndex < dragonCount; dragonIndex++ {
		if strength <= dragons[dragonIndex].Strength {
			gamePassed = false
			break
		}

		strength += dragons[dragonIndex].Bonus
	}

	if gamePassed {
		Fprintf("YES\n")
	} else {
		Fprintf("NO\n")
	}
}

func main() {
	Solution()
	writer.Flush()
}
