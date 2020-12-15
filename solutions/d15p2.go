package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

func D15P2(input []string) string {
	strNums := strings.Split(input[0], ",")
	nums := []int32{}
	for _, s := range strNums {
		n, _ := strconv.Atoi(s)
		nums = append(nums, int32(n))
	}

	var turn int32 = 0
	var num int32 = 0
	var lomemSize int32 = 10000000
	mem := map[int32]int32{}
	lomem := make([]int32, lomemSize)

	for ; turn < int32(len(nums)); turn++ {
		num = nums[turn]
		lomem[num] = turn + 1
	}

	for ; turn != 30000000; turn++ {
		var m int32
		var ok bool
		if num < lomemSize {
			m = lomem[num]
			ok = m != 0
			lomem[num] = turn
		} else {
			m, ok = mem[num]
			mem[num] = turn
		}

		if !ok {
			num = 0
		} else {
			num = turn - m
		}
	}

	return fmt.Sprint(num)
}
