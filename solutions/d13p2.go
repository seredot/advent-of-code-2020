package solutions

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func D13P2(input []string) string {
	lines := strings.Split(input[1], ",")

	type Bus struct {
		num   int64
		index int64
	}

	busses := []Bus{}

	for i, line := range lines {
		num, _ := strconv.Atoi(line)
		if num != 0 {
			busses = append(busses, Bus{
				int64(num), int64(i),
			})
		}
	}

	sort.Slice(busses, func(i, j int) bool {
		return busses[i].num > busses[j].num
	})

	var x int64 = 1
	var inc int64 = 1
	ind := 0

	for {
		x += inc

		if (x+busses[ind].index)%busses[ind].num == 0 {
			inc *= busses[ind].num
			ind++
		}

		if ind > len(busses)-1 {
			return fmt.Sprint(x)
		}
	}

}
