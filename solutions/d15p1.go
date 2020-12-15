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

	turn := 0
	mem := map[int]int{}
	num := 0

	for ; turn < len(nums); turn++ {
		num = nums[turn]
		mem[num] = turn + 1
	}

	for ; turn != 2020; turn++ {
		m, ok := mem[num]
		mem[num] = turn

		if !ok {
			num = 0
		} else {
			num = turn - m
		}
	}

	return fmt.Sprint(num)
}
