package solutions

import (
	"fmt"
	"math"
	"strconv"
)

func D12P1(input []string) string {
	deg := 0.0
	px := 0.0
	py := 0.0

	for _, line := range input {
		cmd := line[:1]
		val, _ := strconv.ParseFloat(line[1:], 64)

		dx := math.Cos(deg / 180.0 * math.Pi)
		dy := math.Sin(deg / 180.0 * math.Pi)

		switch cmd {
		case "F":
			px += dx * val
			py += dy * val
		case "N":
			py += val
		case "S":
			py -= val
		case "E":
			px += val
		case "W":
			px -= val
		case "R":
			deg -= val
		case "L":
			deg += val
		}
	}

	return fmt.Sprint(math.Round(math.Abs(float64(px) - math.Abs(float64(py)))))
}
