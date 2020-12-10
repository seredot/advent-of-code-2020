package solutions

import (
	"sort"
	"strconv"
)

func D10P1(input []string) int {
	nums := make([]int, len(input)+2)
	max := 0

	for i, line := range input {
		num, _ := strconv.Atoi(line)
		nums[i+1] = num
		if num > max {
			max = num
		}
	}

	nums[len(nums)-1] = max + 3
	sort.Ints(nums)

	diffs := map[int]int{1: 0, 2: 0, 3: 0}

	for i, num := range nums {
		if i+1 == len(nums) {
			break
		}

		diff := nums[i+1] - num
		diffs[diff] = diffs[diff] + 1
	}

	return diffs[1] * diffs[3]
}
