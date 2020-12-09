package main

import (
	"fmt"
	"sort"

	"github.com/seredot/advent-of-code-2020/solutions"
	"github.com/seredot/advent-of-code-2020/utils"
)

type solver func([]string) int

var solvers = map[string]solver{
	"d01p1": solutions.D01P1,
	"d01p2": solutions.D01P2,
	"d02p1": solutions.D02P1,
	"d02p2": solutions.D02P2,
	"d03p1": solutions.D03P1,
	"d03p2": solutions.D03P2,
	"d04p1": solutions.D04P1,
	"d04p2": solutions.D04P2,
	"d05p1": solutions.D05P1,
	"d05p2": solutions.D05P2,
}

func main() {
	keys := []string{}
	for k := range solvers {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		solver := solvers[key]
		dayNum := key[1:3]
		input := utils.ReadFile(fmt.Sprintf("input/%s.txt", dayNum))
		result := solver(input)
		fmt.Printf("%s: %d\n", key, result)
	}
}
