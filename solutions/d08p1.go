package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

func D08P1(input []string) string {
	runLines := map[int]bool{}

	acc := 0
	cur := 0

	for {
		if _, ok := runLines[cur]; ok {
			return fmt.Sprint(acc)
		}

		runLines[cur] = true

		line := input[cur]
		tokens := strings.Split(line, " ")
		ins := tokens[0]
		num, _ := strconv.Atoi(tokens[1])

		if ins == "jmp" {
			cur += num
			continue
		}

		if ins == "acc" {
			acc += num
		}

		cur++
	}

	return ""
}
