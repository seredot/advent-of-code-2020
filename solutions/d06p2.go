package solutions

func countAllYes(answers []string) int {
	seenTimes := map[rune]int{}

	for _, a := range answers {
		for _, q := range a {
			if num, ok := seenTimes[q]; ok {
				seenTimes[q] = num + 1
				continue
			}
			seenTimes[q] = 1
		}
	}

	count := 0
	for _, v := range seenTimes {
		if v == len(answers) {
			count++
		}
	}

	return count
}

func D06P2(input []string) int {
	count := 0

	answers := []string{}
	cur := 0

	for {
		answers = append(answers, input[cur])

		cur++
		if cur == len(input) || input[cur] == "" {
			cur++
			count += countAllYes(answers)
			answers = []string{}
		}
		if cur >= len(input) {
			break
		}
	}

	return count
}
