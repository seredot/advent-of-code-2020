package solutions

func countYesQuestions(chars string) int {
	seen := map[rune]bool{}

	for _, c := range chars {
		seen[c] = true
	}

	return len(seen)
}

func D06P1(input []string) int {
	count := 0

	chars := ""
	cur := 0

	for {
		chars += input[cur]

		cur++
		if cur == len(input) || input[cur] == "" {
			count += countYesQuestions(chars)
			chars = ""
		}
		if cur == len(input) {
			break
		}
	}

	return count
}
