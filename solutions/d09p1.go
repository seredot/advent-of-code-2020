package solutions

import "strconv"

func D09P1(input []string) int {
	nums := map[int]int{}

NEXT:
	for i, line := range input {
		num, _ := strconv.Atoi(line)
		if i < 25 {
			nums[i] = num
			continue
		}

		for _, a := range nums {
			b := num - a

			if b == a {
				continue
			}

			for _, c := range nums {
				if b == c {
					nums[i%25] = num
					continue NEXT
				}
			}
		}

		return num
	}

	return 0
}
