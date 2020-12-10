package solutions

import (
	"fmt"
	"strconv"
)

func D01P1(input []string) string {
	hashMap := make(map[int]int)
	for _, line := range input {
		num, _ := strconv.Atoi(line)
		hashMap[num] = num
		other := 2020 - num

		if _, ok := hashMap[other]; ok {
			return fmt.Sprint(num * other)
		}
	}
	return ""
}
