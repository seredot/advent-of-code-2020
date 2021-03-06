package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

func D18P2(input []string) string {
	sum := 0

	var calc func(tokens []string, offset int) (ret int, next int)
	calc = func(tokens []string, offset int) (ret int, next int) {
		ret = 0
		op := ""

		for {
			if offset >= len(tokens) {
				return ret, offset
			}

			t := tokens[offset]

			if t == ")" {
				return ret, offset
			}

			if t == "+" {
				op = t
				offset++
				continue
			}

			var num int
			if t == "(" {
				num, offset = calc(tokens, offset+1)
			} else if t == "*" {
				k, o := calc(tokens, offset+1)
				ret *= k
				offset = o
				return ret, offset
			} else {
				num, _ = strconv.Atoi(t)
			}

			if op == "+" {
				ret = ret + num
				offset++
				continue
			}

			if op == "*" {
				ret = ret * num
				offset++
				continue
			}

			ret = num
			offset++
		}
	}

	for _, line := range input {
		line = strings.ReplaceAll(line, "(", "( ")
		line = strings.ReplaceAll(line, ")", " )")
		tokens := strings.Split(line, " ")
		ret, _ := calc(tokens, 0)
		sum += ret
	}

	return fmt.Sprint(sum)
}
