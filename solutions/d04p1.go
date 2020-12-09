package solutions

import (
	"strings"
)

func validate(props []string) int {
	hashMap := map[string]string{}

	for _, prop := range props {
		tokens := strings.Split(prop, ":")
		if len(tokens) < 2 {
			continue
		}
		k := tokens[0]
		v := tokens[1]
		hashMap[k] = v
	}

	ok := (len(hashMap["byr"]) > 0 &&
		len(hashMap["iyr"]) > 0 &&
		len(hashMap["eyr"]) > 0 &&
		len(hashMap["hgt"]) > 0 &&
		len(hashMap["hcl"]) > 0 &&
		len(hashMap["ecl"]) > 0 &&
		len(hashMap["pid"]) > 0)

	if ok {
		return 1
	}

	return 0
}

func D04P1(input []string) int {
	count := 0

	props := []string{}
	cur := 0

	for {
		props = append(props, strings.Split(input[cur], " ")...)

		cur++
		if cur == len(input) || input[cur] == "" {
			count += validate(props)
			props = []string{}
		}
		if cur == len(input) {
			break
		}
	}

	return count
}
