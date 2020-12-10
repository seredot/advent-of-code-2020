package solutions

import (
	"fmt"
	"strconv"
)

func D01P2(input []string) string {
	data := []int{}

	for _, line := range input {
		num, _ := strconv.Atoi(line)
		data = append(data, num)
	}

	hashMap := make(map[int]int)
	for ind1, num1 := range data {
		left := 2020 - num1

		for ind2, num2 := range data {
			if ind1 == ind2 {
				continue
			}
			num3 := left - num2

			if _, ok := hashMap[num3]; ok {
				return fmt.Sprint(num1 * num2 * num3)
			}

			hashMap[num2] = num2
		}

		hashMap[num1] = num1
	}

	return ""
}
