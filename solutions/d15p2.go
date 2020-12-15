package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

func D15P2(input []string) string {
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
	mem := map[int]int{}
	num := 0

	for {
		turn++

		if turn == 30000000 {
			break
		}

		if turn-1 < len(nums) {
			num = nums[turn-1]
			mem[num] = turn

			if turn < 3 {
				continue
			}
		}

		m, ok := mem[num]
		if !ok {
			m = turn
			mem[num] = m
			num = 0
		} else {
			mem[num] = turn
			num = turn - m
		}

	}

	return fmt.Sprint(num)
}
