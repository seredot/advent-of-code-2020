package solutions

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func D10P2(input []string) string {
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

	count := 0
	clustersOf1 := map[int]int{}

	for i, num := range nums {
		if i+1 == len(nums) {
			break
		}

		diff := nums[i+1] - num
		if diff != 1 {
			clustersOf1[count] = clustersOf1[count] + 1
			count = 0
			continue
		}
		count++
	}

	//	fmt.Println(clustersOf1)

	/* Clusters of 1

	1 (1 combination)

	11
	2 (2 combinations)

	111
	12
	21
	3 (4 combinations)

	1111
	112
	121
	13
	211
	22
	31 (7 combinations)
	*/

	result := int64(math.Pow(2, float64(clustersOf1[2])) * math.Pow(4, float64(clustersOf1[3])) * math.Pow(7, float64(clustersOf1[4])))

	return fmt.Sprint(result)
}
