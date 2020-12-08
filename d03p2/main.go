package main

import (
	"fmt"

	"github.com/seredot/advent-of-code-2020/utils"
)

func collide(lines []string, dx, dy int) int {
	count := 0
	x := 0
	y := 0
	w := len(lines[0])
	h := len(lines)

	for {
		if string(lines[y][x%w]) == "#" {
			count++
		}

		x += dx
		y += dy

		if y >= h {
			break
		}
	}

	return count
}

func main() {
	lines := utils.ReadFile("input.txt")

	product := 1

	product *= collide(lines, 1, 1)
	product *= collide(lines, 3, 1)
	product *= collide(lines, 5, 1)
	product *= collide(lines, 7, 1)
	product *= collide(lines, 1, 2)

	fmt.Println(product)
}
