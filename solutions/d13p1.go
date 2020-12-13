package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

func D13P1(input []string) string {
	t, _ := strconv.Atoi(input[0])
	busses := strings.Split(input[1], ",")

	mywait := 1000000
	mybus := 0
	for _, bus := range busses {
		if bus == "x" {
			continue
		}

		num, _ := strconv.Atoi(bus)
		wait := num - (t % num)

		if wait < mywait {
			mywait = wait
			mybus = num
		}
	}

	return fmt.Sprint(mywait * mybus)
}
