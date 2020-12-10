package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

func D07P2(input []string) string {
	type bag struct {
		children  map[string]int
		contained bool
		countBags int
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
		children := map[string]int{}

		for i := 1; i < len(tokens); i++ {
			count, _ := strconv.Atoi(tokens[i][:1])
			children[tokens[i][2:]] = count
		}

		bags[tokens[0]] = &bag{children, false, 0}
	}

	theOne := bags["shiny gold"]
	theOne.contained = true

	for {
		continueLoop := false

		for _, parent := range bags {
			if parent.contained {
				total := 0
				for k, num := range parent.children {
					child := bags[k]
					if !child.contained {
						continueLoop = true
					}
					child.contained = true
					total += num * (child.countBags + 1)
				}
				if parent.countBags != total {
					parent.countBags = total
					continueLoop = true
				}
			}
		}

		if !continueLoop {
			break
		}
	}

	return fmt.Sprint(theOne.countBags)
}
