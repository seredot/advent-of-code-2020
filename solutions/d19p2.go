package solutions

import (
	"fmt"
	"strings"
)

func D19P2(input []string) string {
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

		if k == "8" {
			tokens[1] = "42 | 42 8"
		}

		if k == "11" {
			tokens[1] = "42 31 | 42 11 31"
		}

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

	seen := map[string]bool{}

	var iter func(key string, then []string, search string, index int)
	iter = func(key string, then []string, search string, index int) {
		rule := rules[key]

		if rule.val != "" {
			if index >= len(search) {
				return
			}

			if rule.val == search[index:index+1] {
				if len(then) > 0 {
					iter(then[0], then[1:], search, index+1)
				} else {
					if index == len(search)-1 {
						seen[search] = true
					}
				}
			}

			return
		}

		g1 := rule.groups[0]
		iter(g1[0], append(g1[1:], then...), search, index)

		if len(rule.groups) > 1 {
			g2 := rule.groups[1]
			iter(g2[0], append(g2[1:], then...), search, index)
		}
	}

	count := 0

	for ; i < len(input); i++ {
		iter(rules["0"].groups[0][0], rules["0"].groups[0][1:], input[i], 0)
		_, ok := seen[input[i]]
		if ok {
			count++
		}
	}

	return fmt.Sprint(count)
}
