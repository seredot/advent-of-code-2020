package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

func D14P1(input []string) string {
	mem := map[int64]int64{}
	var andMask int64 = 0
	var orMask int64 = 0

	for _, line := range input {
		if strings.HasPrefix(line, "mask") {
			andMask = 0
			orMask = 0
			for i, c := range line[7:] {
				if c == '1' {
					orMask |= (1 << (35 - i))
				}

				if c == '0' {
					andMask |= (1 << (35 - i))
				}
			}

			andMask = ^andMask
			continue
		}

		parts := strings.Split(line, " ")
		addr, _ := strconv.ParseInt(parts[0][4:len(parts[0])-1], 10, 64)
		val, _ := strconv.ParseInt(parts[len(parts)-1], 10, 64)

		val &= andMask
		val |= orMask

		mem[addr] = val
	}

	var acc int64 = 0
	for _, v := range mem {
		acc += v
	}

	return fmt.Sprint(acc)
}
