package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/seredot/advent-of-code-2020/utils"
)

func validByr(val string) bool {
	if len(val) != 4 {
		return false
	}

	yr, _ := strconv.Atoi(val)

	return !(yr < 1920 || yr > 2002)
}

func validIyr(val string) bool {
	if len(val) != 4 {
		return false
	}

	yr, _ := strconv.Atoi(val)

	return !(yr < 2010 || yr > 2020)
}

func validEyr(val string) bool {
	if len(val) != 4 {
		return false
	}

	yr, _ := strconv.Atoi(val)

	return !(yr < 2020 || yr > 2030)
}

var reHeightCm = regexp.MustCompile(`[0-9]{3}cm`)
var reHeightIn = regexp.MustCompile(`[0-9]{2}in`)

func validHgt(val string) bool {
	if reHeightCm.MatchString(val) {
		cm, _ := strconv.Atoi(val[:3])
		return !(cm < 150 || cm > 193)
	}

	if reHeightIn.MatchString(val) {
		in, _ := strconv.Atoi(val[:2])
		return !(in < 59 || in > 76)
	}

	return false
}

var reHairColor = regexp.MustCompile(`^#([a-f0-9]{6})$`)

func validHcl(val string) bool {
	return reHairColor.MatchString(val)
}

var eyeColors = map[string]string{"amb": "", "blu": "", "brn": "", "gry": "", "grn": "", "hzl": "", "oth": ""}

func validEcl(val string) bool {
	_, ok := eyeColors[val]
	return ok
}

var rePid = regexp.MustCompile(`^([0-9]{9})$`)

func validPid(val string) bool {
	return rePid.MatchString(val)
}

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

	ok := (validByr(hashMap["byr"]) &&
		validIyr(hashMap["iyr"]) &&
		validEyr(hashMap["eyr"]) &&
		validHgt(hashMap["hgt"]) &&
		validHcl(hashMap["hcl"]) &&
		validEcl(hashMap["ecl"]) &&
		validPid(hashMap["pid"]))

	if ok {
		return 1
	}

	return 0
}

func main() {
	lines := utils.ReadFile("input.txt")

	count := 0

	props := []string{}
	cur := 0

	for {
		props = append(props, strings.Split(lines[cur], " ")...)

		cur++
		if cur == len(lines) || lines[cur] == "" {
			count += validate(props)
			props = []string{}
		}
		if cur == len(lines) {
			break
		}
	}

	fmt.Println(count)
}
