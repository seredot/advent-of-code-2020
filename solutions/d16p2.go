package solutions

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func D16P2(input []string) string {
	type field struct {
		name    string
		min1    int
		max1    int
		min2    int
		max2    int
		indices []int
		index   int
	}

	fields := map[string]*field{}
	cols := [][]int{}

	for _, line := range input[:20] {
		s := strings.Split(line, ":")
		name := s[0]
		r := strings.ReplaceAll(s[1], " or ", ",")
		r = strings.ReplaceAll(r, " ", "")
		rs := strings.Split(r, ",")
		r1 := strings.Split(rs[0], "-")
		r2 := strings.Split(rs[1], "-")
		min1, _ := strconv.Atoi(r1[0])
		max1, _ := strconv.Atoi(r1[1])
		min2, _ := strconv.Atoi(r2[0])
		max2, _ := strconv.Atoi(r2[1])
		fields[name] = &field{name, min1, max1, min2, max2, []int{}, -1}
		cols = append(cols, make([]int, 0, len(input)))
	}

	parseRow := func(row string) {
		tokens := strings.Split(row, ",")

		for i, s := range tokens {
			num, _ := strconv.Atoi(s)
			cols[i] = append(cols[i], num)
		}
	}

	parseRow(input[22])
	stopLines := map[int]bool{0: true, 9: true, 10: true, 11: true, 16: true, 17: true, 18: true, 37: true, 42: true, 44: true, 53: true, 55: true, 56: true, 61: true, 67: true, 85: true, 87: true, 88: true, 98: true, 100: true, 101: true, 105: true, 107: true, 116: true, 118: true, 128: true, 130: true, 134: true, 139: true, 140: true, 144: true, 148: true, 151: true, 157: true, 159: true, 160: true, 163: true, 169: true, 174: true, 175: true, 179: true, 180: true, 193: true, 194: true, 201: true, 203: true, 208: true, 211: true, 212: true, 214: true, 225: true, 229: true, 233: true, 240: true}
	for i, row := range input[25:] {
		if _, ok := stopLines[i]; !ok {
			parseRow(row)
		}
	}

	fieldList := []*field{}

	for _, v := range fields {
		fieldList = append(fieldList, v)

	NEXT_COL:
		for c, col := range cols {
			for _, num := range col {
				if !((num >= v.min1 && num <= v.max1) || (num >= v.min2 && num <= v.max2)) {
					continue NEXT_COL
				}
			}

			v.indices = append(v.indices, c)
		}
	}

	sort.Slice(fieldList, func(i int, j int) bool {
		return len(fieldList[i].indices) < len(fieldList[j].indices)
	})

	matched := map[int]bool{}

	acc := 1

NEXT_FIELD:
	for _, field := range fieldList {
		for _, ind := range field.indices {
			_, ok := matched[ind]

			if !ok {
				matched[ind] = true
				field.index = ind

				if strings.HasPrefix(field.name, "departure") {
					acc *= cols[ind][0]
				}
				continue NEXT_FIELD
			}
		}

		fmt.Printf("Field not found for %s\n", field.name)
	}

	return fmt.Sprint(acc)
}
