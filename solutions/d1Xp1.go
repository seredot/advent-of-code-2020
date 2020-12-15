package solutions

import (
	"fmt"
	"strconv"
)

func D1XP1(input []string) string {
	nums := map[int]int{}

	count := 0

	for _, line := range input {
		num, _ := strconv.Atoi(line)
		nums[num] = num

		// This is a template
	}

	return fmt.Sprint(count)
}
