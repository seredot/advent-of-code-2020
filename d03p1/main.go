package main

import (
	"fmt"

	"github.com/seredot/advent-of-code-2020/utils"
)

func main() {
	lines := utils.ReadFile("input.txt")
	count := 0
	x := 0

	for _, line := range lines {
		if string(line[x%len(line)]) == "#" {
			count++
		}
		x += 3
	}

	fmt.Println(count)
}
