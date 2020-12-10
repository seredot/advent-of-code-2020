package solutions

import "fmt"

func D03P1(input []string) string {
	count := 0
	x := 0

	for _, line := range input {
		if string(line[x%len(line)]) == "#" {
			count++
		}
		x += 3
	}

	return fmt.Sprint(count)
}
