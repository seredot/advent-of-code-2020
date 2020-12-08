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

	pos1, _ := strconv.Atoi(tokens[0])
	pos1--
	pos2, _ := strconv.Atoi(tokens[1])
	pos2--
	letter := tokens[2]
	pwd := tokens[4]

	char1 := ""
	if pos1 < len(pwd) {
		char1 = string(pwd[pos1])
	}

	char2 := ""
	if pos2 < len(pwd) {
		char2 = string(pwd[pos2])
	}

	if (char1 == letter && char2 != letter) || (char2 == letter && char1 != letter) {
		return 1
	}

	return 0
}

func main() {
	lines := readFile("input.txt")
	count := 0

	for _, line := range lines {
		count += checkValid(line)
	}

	fmt.Println(count)
}
