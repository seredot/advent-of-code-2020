package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

func D15P1(input []string) string {
	strNums := strings.Split(input[0], ",")
	nums := []int{}
	for _, s := range strNums {
		n, _ := strconv.Atoi(s)
		nums = append(nums, n)
	}

	type tuple struct {
		p1 int
		p2 int
	}

	turn := 0
	mem := map[int]*tuple{}
	num := 0

	for {
		turn++

		if turn == 2020 {
			break
		}

		if turn-1 < len(nums) {
			num = nums[turn-1]
			mem[num] = &tuple{
				p1: turn,
				p2: turn,
			}

			if turn < 3 {
				continue
			}
		}

		m, ok := mem[num]
		if !ok {
			m = &tuple{
				p1: turn,
				p2: turn,
			}
			mem[num] = m
			num = 0
		} else {
			num = turn - m.p1
			m.p2 = m.p1
			m.p1 = turn
		}

	}

	return fmt.Sprint(num)
}
