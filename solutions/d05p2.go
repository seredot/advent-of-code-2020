package solutions

type seatInfo struct {
	row int
	col int
}

func (s seatInfo) id() int {
	return s.row*8 + s.col
}

func bits(b byte) int {
	bs := string(b)
	if bs == "F" || bs == "L" {
		return 0
	}

	return 1
}

func decode2(code string) seat {
	row := 0
	row |= bits(code[0]) << 6
	row |= bits(code[1]) << 5
	row |= bits(code[2]) << 4
	row |= bits(code[3]) << 3
	row |= bits(code[4]) << 2
	row |= bits(code[5]) << 1
	row |= bits(code[6])
	col := 0
	col |= bits(code[7]) << 2
	col |= bits(code[8]) << 1
	col |= bits(code[9])

	return seat{row, col}
}

func D05P2(input []string) int {
	hashMap := map[int]bool{}

	for _, line := range input {
		s := decode2(line)
		id := s.id()
		hashMap[id] = true
	}

	for i := 1; i < 1000; i++ {
		_, prev := hashMap[i-1]
		_, cur := hashMap[i]
		_, next := hashMap[i+1]
		if prev && !cur && next {
			return i
		}
	}

	return 0
}
