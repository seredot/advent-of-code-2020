package main

import (
	"fmt"

	"github.com/seredot/advent-of-code-2020/utils"
)

type seat struct {
	row int
	col int
}

func (s seat) id() int {
	return s.row*8 + s.col
}

func bit(b byte) int {
	bs := string(b)
	if bs == "F" || bs == "L" {
		return 0
	}

	return 1
}

func decode(code string) seat {
	row := 0
	row |= bit(code[0]) << 6
	row |= bit(code[1]) << 5
	row |= bit(code[2]) << 4
	row |= bit(code[3]) << 3
	row |= bit(code[4]) << 2
	row |= bit(code[5]) << 1
	row |= bit(code[6])
	col := 0
	col |= bit(code[7]) << 2
	col |= bit(code[8]) << 1
	col |= bit(code[9])

	return seat{row, col}
}

func main() {
	lines := utils.ReadFile("input.txt")

	hashMap := map[int]bool{}

	for _, line := range lines {
		s := decode(line)
		id := s.id()
		hashMap[id] = true
	}

	for i := 1; i < 1000; i++ {
		_, prev := hashMap[i-1]
		_, cur := hashMap[i]
		_, next := hashMap[i+1]
		if prev && !cur && next {
			fmt.Println(i)
		}
	}
}
