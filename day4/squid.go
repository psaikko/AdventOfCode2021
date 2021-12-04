package day4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type board [5][5]int

func readBoard(s *bufio.Scanner) (board, bool) {

	b := board{}

	stringsLines := readNStringsLines(s, 5)

	if len(stringsLines) == 5 {
		for ri, row := range stringsLines {
			for ci, val := range row {
				v, err := strconv.Atoi(val)
				if err != nil {
					panic(err)
				}
				b[ri][ci] = v
			}
		}
	} else {
		return b, false
	}

	return b, true
}

// ReadStringsLines returns lines split by whitespace from an io.Reader
func readNStringsLines(s *bufio.Scanner, n int) [][]string {
	lines := make([]string, n)
	stringsLines := make([][]string, 0, n)

	for i := 0; i < n; i++ {
		if !s.Scan() {
			return stringsLines
		}
		line := s.Text()
		lines[i] = strings.TrimSpace(line)
	}

	for _, line := range lines {
		stringsLines = append(stringsLines, strings.Fields(line))
	}

	return stringsLines
}

func (b *board) mark(number int) bool {
	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			if b[r][c] == number {
				b[r][c] = -1
				return true
			}
		}
	}
	return false
}

func (b board) isBingo() bool {
	for r := 0; r < 5; r++ {
		fullRow := true
		for c := 0; c < 5; c++ {
			if b[r][c] >= 0 {
				fullRow = false
			}
		}
		if fullRow {
			return true
		}
	}

	for c := 0; c < 5; c++ {
		fullCol := true
		for r := 0; r < 5; r++ {
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
	for c := 0; c < 5; c++ {
		for r := 0; r < 5; r++ {
			if b[r][c] >= 0 {
				sum += b[r][c]
			}
		}
	}
	return sum
}

func Run() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	numbersLine := scanner.Text()
	numbersStrings := strings.Split(numbersLine, ",")
	numbersInts := make([]int, len(numbersStrings))
	for i, s := range numbersStrings {
		v, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		numbersInts[i] = v
	}

	boards := make([]board, 0, 3)

	for {
		scanner.Scan()
		b, ok := readBoard(scanner)
		if !ok {
			break
		}
		boards = append(boards, b)
	}

	winners := 0
	won := make([]bool, len(boards))

	for _, n := range numbersInts {
		for i := range boards {
			boards[i].mark(n)
			if !won[i] && boards[i].isBingo() {
				if winners == 0 || winners == len(boards)-1 {
					fmt.Println(boards[i])
					fmt.Println(n)
					fmt.Println(n * boards[i].sum())
				}
				won[i] = true
				winners++
			}
		}
	}
}
