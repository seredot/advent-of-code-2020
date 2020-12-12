package solutions

import (
	"fmt"
	"math"
	"strconv"
)

func D12P2(input []string) string {
	px := 0
	py := 0
	dx := 10
	dy := 1

	rotate := func(deg int) {
		t := dx

		switch deg {
		case 90:
			dx = -dy
			dy = t
		case 180:
			dx = -dx
			dy = -dy
		case 270:
			dx = dy
			dy = -t
		}
	}

	for _, line := range input {
		cmd := line[:1]
		v, _ := strconv.ParseInt(line[1:], 0, 32)
		val := int(v)

		switch cmd {
		case "F":
			px += dx * val
			py += dy * val
		case "N":
			dy += val
		case "S":
			dy -= val
		case "E":
			dx += val
		case "W":
			dx -= val
		case "R":
			rotate(360 - val)
		case "L":
			rotate(val)
		}
	}

	return fmt.Sprint(math.Abs(float64(px)) + math.Abs(float64(py)))
}
