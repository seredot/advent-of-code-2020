package solutions

type seat struct {
	row int
	col int
}

func (s seat) id() int {
	return s.row*8 + s.col
}

func bit(b byte) int {
	bs := string(b)
	if bs == "F" || bs == "L" {
		return 0
	}

	return 1
}

func decode(code string) seat {
	row := 0
	row |= bit(code[0]) << 6
	row |= bit(code[1]) << 5
	row |= bit(code[2]) << 4
	row |= bit(code[3]) << 3
	row |= bit(code[4]) << 2
	row |= bit(code[5]) << 1
	row |= bit(code[6])
	col := 0
	col |= bit(code[7]) << 2
	col |= bit(code[8]) << 1
	col |= bit(code[9])

	return seat{row, col}
}

func D05P1(input []string) int {
	max := -1

	for _, line := range input {
		s := decode(line)
		id := s.id()
		if id > max {
			max = id
		}
	}

	return max
}
