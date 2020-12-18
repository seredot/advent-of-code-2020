package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/seredot/advent-of-code-2020/solutions"
	"github.com/seredot/advent-of-code-2020/utils"
)

type solver func([]string) string

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
	"d06p1": solutions.D06P1,
	"d06p2": solutions.D06P2,
	"d07p1": solutions.D07P1,
	"d07p2": solutions.D07P2,
	"d08p1": solutions.D08P1,
	"d08p2": solutions.D08P2,
	"d09p1": solutions.D09P1,
	"d09p2": solutions.D09P2,
	"d10p1": solutions.D10P1,
	"d10p2": solutions.D10P2,
	"d11p1": solutions.D11P1,
	"d11p2": solutions.D11P2,
	"d12p1": solutions.D12P1,
	"d12p2": solutions.D12P2,
	"d13p1": solutions.D13P1,
	"d13p2": solutions.D13P2,
	"d14p1": solutions.D14P1,
	"d14p2": solutions.D14P2,
	"d15p1": solutions.D15P1,
	"d15p2": solutions.D15P2,
	"d16p1": solutions.D16P1,
	"d16p2": solutions.D16P2,
	"d17p1": solutions.D17P1,
	"d17p2": solutions.D17P2,
	"d18p1": solutions.D18P1,
	"d18p2": solutions.D18P2,
	/*
		"d19p1": solutions.D19P1,
		"d19p2": solutions.D19P2,
		"d20p1": solutions.D20P1,
		"d20p2": solutions.D20P2,
		"d21p1": solutions.D21P1,
		"d21p2": solutions.D21P2,
		"d22p1": solutions.D22P1,
		"d22p2": solutions.D22P2,
		"d23p1": solutions.D23P1,
		"d23p2": solutions.D23P2,
		"d24p1": solutions.D24P1,
		"d24p2": solutions.D24P2,
		"d25p1": solutions.D25P1,
		"d25p2": solutions.D25P2,
	*/
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
		startTime := time.Now()
		result := solver(input)
		timeSpan := time.Now().Sub(startTime)
		timeRound := timeSpan.Round(time.Millisecond)
		if timeSpan < time.Millisecond {
			timeRound = timeSpan.Round(time.Microsecond)
		}
		fmt.Printf("%s (%v)\t%s\n", key, timeRound, result)
	}
}
