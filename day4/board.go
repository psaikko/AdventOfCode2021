package day4

import (
	"bufio"
	"strconv"
	"strings"
)

const boardSize = 5

type board [boardSize][boardSize]int

// scanNStringsLines returns n lines of strings split by strings.Fields
func scanNStringsLines(s *bufio.Scanner, n int) [][]string {
	lines := make([]string, n)
	stringsLines := make([][]string, 0, n)

	for i := 0; i < n; i++ {
		if !s.Scan() {
			return stringsLines
		}
		lines[i] = strings.TrimSpace(s.Text())
	}

	for _, line := range lines {
		stringsLines = append(stringsLines, strings.Fields(line))
	}

	return stringsLines
}

func readBoard(s *bufio.Scanner) (board, bool) {
	b := board{}
	stringsLines := scanNStringsLines(s, boardSize)

	if len(stringsLines) == boardSize {
		for ri, row := range stringsLines {
			for ci, val := range row {
				v, err := strconv.Atoi(val)
				if err != nil {
					panic(err)
				}
				b[ri][ci] = v
			}
		}
		return b, true
	}
	return b, false
}

func (b *board) mark(number int) bool {
	for r := 0; r < boardSize; r++ {
		for c := 0; c < boardSize; c++ {
			if b[r][c] == number {
				b[r][c] = -1
				return true
			}
		}
	}
	return false
}

func (b board) isBingo() bool {
	for r := 0; r < boardSize; r++ {
		fullRow := true
		for c := 0; c < boardSize; c++ {
			if b[r][c] >= 0 {
				fullRow = false
			}
		}
		if fullRow {
			return true
		}
	}

	for c := 0; c < boardSize; c++ {
		fullCol := true
		for r := 0; r < boardSize; r++ {
			if b[r][c] >= 0 {
				fullCol = false
			}
		}
		if fullCol {
			return true
		}
	}

	return false
}

func (b board) sum() int {
	sum := 0
	for c := 0; c < boardSize; c++ {
		for r := 0; r < boardSize; r++ {
			if b[r][c] >= 0 {
				sum += b[r][c]
			}
		}
	}
	return sum
}
