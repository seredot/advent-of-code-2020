package solutions

import (
	"math"
	"strconv"
)

func D09P2(input []string) int {
	numbers := make([]int, len(input))
	for i := 0; i < len(input); i++ {
		n, _ := strconv.Atoi(input[i])
		numbers[i] = n
	}

	theNum := 248131121 // The magic number to be searched.
	end := 590          // The end inde of the magic number.
	start := end - 1

	for {
		sum := 0
		for i := start; i <= end; i++ {
			sum += numbers[i]
		}

		if sum == theNum {
			break
		}

		if sum > theNum {
			end--
		}

		if sum < theNum {
			start--
		}
	}

	min := math.MaxInt32
	max := math.MinInt32
	for i := start; i <= end; i++ {
		if numbers[i] < min {
			min = numbers[i]
		}

		if numbers[i] > max {
			max = numbers[i]
		}
	}

	return min + max
}
