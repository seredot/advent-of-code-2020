package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

func D16P1(input []string) string {
	type ranges struct {
		min1 int
		max1 int
		min2 int
		max2 int
	}

	fields := map[string]ranges{}

	for _, line := range input[:20] {
		s := strings.Split(line, ":")
		r := strings.ReplaceAll(s[1], " or ", ",")
		r = strings.ReplaceAll(r, " ", "")
		rs := strings.Split(r, ",")
		r1 := strings.Split(rs[0], "-")
		r2 := strings.Split(rs[1], "-")
		min1, _ := strconv.Atoi(r1[0])
		max1, _ := strconv.Atoi(r1[1])
		min2, _ := strconv.Atoi(r2[0])
		max2, _ := strconv.Atoi(r2[1])
		fields[s[0]] = ranges{min1, max1, min2, max2}
	}

	acc := 0
	for _, line := range input[25:] {
		tokens := strings.Split(line, ",")

		for _, s := range tokens {
			any := false
			num, _ := strconv.Atoi(s)

			for _, v := range fields {
				if (num >= v.min1 && num <= v.max1) || (num >= v.min2 && num <= v.max2) {
					any = true
					break
				}
			}

			if !any {
				acc += num
				//fmt.Println(_) // prints corrupt lines used as input in p2
			}
		}
	}

	return fmt.Sprint(acc)
}
