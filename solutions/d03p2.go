package solutions

func collide(lines []string, dx, dy int) int {
	count := 0
	x := 0
	y := 0
	w := len(lines[0])
	h := len(lines)

	for {
		if string(lines[y][x%w]) == "#" {
			count++
		}

		x += dx
		y += dy

		if y >= h {
			break
		}
	}

	return count
}

func D03P2(input []string) int {
	product := 1

	product *= collide(input, 1, 1)
	product *= collide(input, 3, 1)
	product *= collide(input, 5, 1)
	product *= collide(input, 7, 1)
	product *= collide(input, 1, 2)

	return product
}
