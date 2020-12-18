package solutions

import (
	"fmt"
)

func D17P2(input []string) string {
	var mem *map[int]byte
	var next *map[int]byte

	set := func(x, y, z, w int, v byte) {
		if v == 1 {
			for dx := -1; dx < 2; dx++ {
				for dy := -1; dy < 2; dy++ {
					for dz := -1; dz < 2; dz++ {
						for dw := -1; dw < 2; dw++ {
							addr := (x+dx)<<0 | (y+dy)<<8 | (z+dz)<<16 | (w+dw)<<24
							(*next)[addr] = (*next)[addr] + 0
						}
					}
				}
			}
		}

		addr := x<<0 | y<<8 | z<<16 | w<<24
		(*next)[addr] = v
	}

	get := func(x, y, z, w int) byte {
		addr := x<<0 | y<<8 | z<<16 | w<<24
		return (*mem)[addr]
	}

	next = &map[int]byte{}
	for y, line := range input {
		for x, c := range line {
			if c == '#' {
				x0 := x + 1<<3
				y0 := y + 1<<3
				z0 := 0 + 1<<3
				w0 := 0 + 1<<3
				set(x0, y0, z0, w0, 1)
			}
		}
	}

	count := func(x, y, z, w int) int {
		c := 0

		for dx := -1; dx < 2; dx++ {
			for dy := -1; dy < 2; dy++ {
				for dz := -1; dz < 2; dz++ {
					for dw := -1; dw < 2; dw++ {
						c += int(get(x+dx, y+dy, z+dz, w+dw))
					}
				}
			}
		}

		return c - int(get(x, y, z, w))
	}

	mask := 0b11111111
	iter := func() {
		for k, v := range *mem {
			x := k & mask
			y := k >> 8 & mask
			z := k >> 16 & mask
			w := k >> 24 & mask
			c := count(x, y, z, w)

			if v == 0 && c == 3 {
				set(x, y, z, w, 1)
				continue
			}

			if v == 1 {
				if c == 2 || c == 3 {
					set(x, y, z, w, 1)
				} else {
					set(x, y, z, w, 0)
				}
			}
		}
	}

	for i := 0; i < 6; i++ {
		mem = next
		next = &map[int]byte{}
		iter()
	}

	acc := 0
	for _, v := range *next {
		acc += int(v)
	}

	return fmt.Sprint(acc)
}
