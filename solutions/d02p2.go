package solutions

import (
	"strconv"
	"strings"
)

func checkValidity(line string) int {
	replaced := strings.ReplaceAll(line, "-", " ")
	replaced = strings.ReplaceAll(replaced, ":", " ")
	tokens := strings.Split(replaced, " ")

	pos1, _ := strconv.Atoi(tokens[0])
	pos1--
	pos2, _ := strconv.Atoi(tokens[1])
	pos2--
	letter := tokens[2]
	pwd := tokens[4]

	char1 := ""
	if pos1 < len(pwd) {
		char1 = string(pwd[pos1])
	}

	char2 := ""
	if pos2 < len(pwd) {
		char2 = string(pwd[pos2])
	}

	if (char1 == letter && char2 != letter) || (char2 == letter && char1 != letter) {
		return 1
	}

	return 0
}

func D02P2(input []string) int {
	count := 0

	for _, line := range input {
		count += checkValidity(line)
	}

	return count
}
