package day4

import (
	"adventofcode/common"
	"bufio"
	"fmt"
	"os"
)

func Run() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	bingoNumbers := common.ScanIntLine(scanner, ",")

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

	for _, n := range bingoNumbers {
		for i := range boards {
			boards[i].mark(n)
			if !won[i] && boards[i].isBingo() {
				if winners == 0 {
					fmt.Println("First winner:", boards[i])
					fmt.Println(n * boards[i].sum())
				}

				if winners == len(boards)-1 {
					fmt.Println("Last winner:", boards[i])
					fmt.Println(n * boards[i].sum())
				}

				won[i] = true
				winners++
			}
		}
	}
}
