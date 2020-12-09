package solutions

import (
	"strconv"
	"strings"
)

func checkValid(line string) int {
	replaced := strings.ReplaceAll(line, "-", " ")
	replaced = strings.ReplaceAll(replaced, ":", " ")
	tokens := strings.Split(replaced, " ")

	min, _ := strconv.Atoi(tokens[0])
	max, _ := strconv.Atoi(tokens[1])
	letter := tokens[2]
	pwd := tokens[4]

	count := 0

	for _, char := range pwd {
		if string(char) == letter {
			count++
		}

		if count > max {
			return 0
		}
	}

	if count < min {
		return 0
	}

	return 1
}

func D02P1(input []string) int {
	count := 0

	for _, line := range input {
		count += checkValid(line)
	}

	return count
}
