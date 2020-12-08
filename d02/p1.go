package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFile(path string) []string {
	ret := []string{}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return ret
}

func checkValid(line string) int {
	replaced := strings.ReplaceAll(line, "-", " ")
	replaced = strings.ReplaceAll(replaced, ":", " ")
	tokens := strings.Split(replaced, " ")

	min, _ := strconv.Atoi(tokens[0])
	max, _ := strconv.Atoi(tokens[1])
	letter := tokens[2]
	pwd := tokens[4]

	count := 0

	for _, char := range pwd {
		if string(char) == letter {
			count++
		}

		if count > max {
			return 0
		}
	}

	if count < min {
		return 0
	}

	return 1
}

func main() {
	lines := readFile("input.txt")
	count := 0

	for _, line := range lines {
		count += checkValid(line)
	}

	fmt.Println(count)
}
