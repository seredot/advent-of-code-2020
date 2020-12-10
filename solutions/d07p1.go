package solutions

import (
	"fmt"
	"strings"
)

func D07P1(input []string) string {
	type bag struct {
		children map[string]bool
		contains bool
	}

	bags := map[string]*bag{}

	for _, line := range input {
		cleanLine := strings.ReplaceAll(line, " contain no other bags.", "")
		cleanLine = strings.ReplaceAll(cleanLine, " contain", ",")
		cleanLine = strings.ReplaceAll(cleanLine, " bags", "")
		cleanLine = strings.ReplaceAll(cleanLine, " bag", "")
		cleanLine = strings.ReplaceAll(cleanLine, ", ", ",")
		cleanLine = strings.ReplaceAll(cleanLine, ".", "")

		tokens := strings.Split(cleanLine, ",")
		children := map[string]bool{}

		for i := 1; i < len(tokens); i++ {
			children[tokens[i][2:]] = true
		}

		bags[tokens[0]] = &bag{children, false}
	}

	term := "shiny gold"

	for {
		foundOne := false

		for _, v := range bags {
			if v.contains {
				continue
			}

			if _, ok := v.children[term]; ok {
				v.contains = true
				foundOne = true
				continue
			}

			for k := range v.children {
				if bags[k].contains {
					v.contains = true
					foundOne = true
				}
			}
		}

		if !foundOne {
			break
		}
	}

	count := 0

	for _, b := range bags {
		if b.contains {
			count++
		}
	}

	return fmt.Sprint(count)
}
