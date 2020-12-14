package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

func D14P2(input []string) string {
	mem := map[int64]int64{}
	var addrMask int64 = 0
	var orMask int64 = 0

	var setMem func(addr, mask, val int64, offset int)
	setMem = func(addr, mask, val int64, offset int) {
		for i := offset; i < 36; i++ {
			if (mask>>i)&1 == 1 {
				setMem(addr|(1<<i), mask, val, i+1)
				setMem(addr&(^(1 << i)), mask, val, i+1)
				return
			}
		}
		mem[addr] = val
	}

	for _, line := range input {
		if strings.HasPrefix(line, "mask") {
			addrMask = 0
			orMask = 0
			for i, c := range line[7:] {
				if c == '1' {
					orMask |= (1 << (35 - i))
				}

				if c == 'X' {
					addrMask |= (1 << (35 - i))
				}
			}

			continue
		}

		parts := strings.Split(line, " ")
		addr, _ := strconv.ParseInt(parts[0][4:len(parts[0])-1], 10, 64)
		val, _ := strconv.ParseInt(parts[len(parts)-1], 10, 64)

		addr |= orMask

		setMem(addr, addrMask, val, 0)
	}

	var acc int64 = 0
	for _, v := range mem {
		acc += v
	}

	return fmt.Sprint(acc)
}
