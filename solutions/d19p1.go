package solutions

import (
	"fmt"
	"strings"
)

func D19P1(input []string) string {
	type rule struct {
		val    string
		groups [][]string
	}

	rules := map[string]rule{}

	i := 0

	for ; ; i++ {
		line := input[i]

		if line == "" {
			i++
			break
		}

		tokens := strings.Split(line, ": ")
		k := tokens[0]

		if strings.HasPrefix(tokens[1], "\"") {
			rules[k] = rule{
				val: strings.Trim(tokens[1], "\""),
			}

			continue
		}

		r := rule{
			groups: [][]string{},
		}

		for _, str := range strings.Split(tokens[1], " | ") {
			r.groups = append(r.groups, strings.Split(str, " "))
		}

		rules[k] = r
	}

	var iter func(key, search string, index int) (bool, int)
	iter = func(key, search string, index int) (bool, int) {
		rule := rules[key]

		if index == len(search) {
			return true, 0
		}

		if rule.val != "" {
			if search[index:index+1] == rule.val {
				return true, index + 1
			}

			return false, 0
		}

		offset := 0

		for _, g := range rule.groups {
			groupResult := true

			offset = index

			for _, n := range g {
				result, next := iter(n, search, offset)
				if !result {
					groupResult = false
					break
				}

				offset = next
			}

			if groupResult {
				return true, offset
			}
		}

		return false, index
	}

	count := 0

	for ; i < len(input); i++ {
		ok, offset := iter("0", input[i], 0)
		if ok && offset == len(input[i]) {
			count++
		}
	}

	return fmt.Sprint(count)
}
