package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

func D08P2(input []string) string {
	type lines = map[int]bool

	interpret := func() (acc int, runLines lines, success bool) {
		runLines = lines{}

		acc = 0
		cur := 0

		for {
			if cur == len(input) {
				success = true
				return
			}

			if _, ok := runLines[cur]; ok {
				success = false
				return
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
	}

	_, loopingLines, _ := interpret()

	for cur := range loopingLines {
		line := input[cur]
		if strings.HasPrefix(line, "acc") {
			continue
		}

		temp := line
		if strings.HasPrefix(line, "jmp") {
			input[cur] = strings.Replace(line, "jmp", "nop", 1)
		} else {
			input[cur] = strings.Replace(line, "nop", "jmp", 1)
		}

		acc, _, success := interpret()
		if success {
			return fmt.Sprint(acc)
		}

		input[cur] = temp
	}

	return ""
}
