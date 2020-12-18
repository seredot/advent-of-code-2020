package solutions

import (
	"fmt"
)

func D17P1(input []string) string {
	var mem *map[int]byte
	var next *map[int]byte

	set := func(x, y, z int, v byte) {
		if v == 1 {
			for dx := -1; dx < 2; dx++ {
				for dy := -1; dy < 2; dy++ {
					for dz := -1; dz < 2; dz++ {
						addr := (x+dx)<<0 | (y+dy)<<10 | (z+dz)<<20
						(*next)[addr] = (*next)[addr] + 0
					}
				}
			}
		}

		addr := x<<0 | y<<10 | z<<20
		(*next)[addr] = v
	}

	get := func(x, y, z int) byte {
		addr := x<<0 | y<<10 | z<<20
		return (*mem)[addr]
	}

	next = &map[int]byte{}
	for y, line := range input {
		for x, c := range line {
			if c == '#' {
				x0 := x + 1<<4
				y0 := y + 1<<4
				z0 := 0 + 1<<4
				set(x0, y0, z0, 1)
			}
		}
	}

	count := func(x, y, z int) int {
		c := 0

		for dx := -1; dx < 2; dx++ {
			for dy := -1; dy < 2; dy++ {
				for dz := -1; dz < 2; dz++ {
					c += int(get(x+dx, y+dy, z+dz))
				}
			}
		}

		return c - int(get(x, y, z))
	}

	mask := 0b1111111111
	iter := func() {
		for k, v := range *mem {
			x := k & mask
			y := k >> 10 & mask
			z := k >> 20 & mask
			c := count(x, y, z)

			if v == 0 && c == 3 {
				set(x, y, z, 1)
				continue
			}

			if v == 1 {
				if c == 2 || c == 3 {
					set(x, y, z, 1)
				} else {
					set(x, y, z, 0)
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
