package solutions

import (
	"fmt"
)

func D11P2(input []string) string {
	w := len(input[0])
	h := len(input)
	size := w * h
	state := make([]rune, size)

	for y, line := range input {
		for x, r := range line {
			state[y*w+x] = r
		}
	}

	occupied := func(y, x, vecY, vecX int) int {
		for {
			y += vecY
			x += vecX

			if y < 0 || x < 0 || y >= h || x >= w {
				return 0
			}

			if state[y*w+x] == '#' {
				return 1
			}

			if state[y*w+x] == 'L' {
				return 0
			}
		}
	}

	iterate := func() bool {
		next := make([]rune, size)
		hasChange := false

		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				if state[y*w+x] == '.' {
					next[y*w+x] = '.'
					continue
				}

				count := 0
				count += occupied(y, x, -1, -1)
				count += occupied(y, x, -1, 0)
				count += occupied(y, x, -1, +1)
				count += occupied(y, x, 0, -1)
				count += occupied(y, x, 0, +1)
				count += occupied(y, x, +1, -1)
				count += occupied(y, x, +1, 0)
				count += occupied(y, x, +1, +1)

				if state[y*w+x] == 'L' && count == 0 {
					next[y*w+x] = '#'
					hasChange = true
					continue
				}

				if state[y*w+x] == '#' && count >= 5 {
					next[y*w+x] = 'L'
					hasChange = true
					continue
				}

				next[y*w+x] = state[y*w+x]
			}
		}

		state = next
		return hasChange
	}

	for {
		hasChange := iterate()
		if !hasChange {
			break
		}
	}

	count := 0
	for _, r := range state {
		if r == '#' {
			count++
		}
	}

	return fmt.Sprint(count)
}
