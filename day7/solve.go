package day7

import (
	"adventofcode/common"
	"bufio"
	"fmt"
	"os"
)

func Run() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	positions := common.ScanIntLine(scanner, ",")
	maxPos := common.Max(positions...)

	bestScore1 := 9999999999
	bestScore2 := 9999999999

	moveCosts := make([]int, maxPos+1)
	moveCosts[1] = 1
	for i := 1; i < maxPos; i += 1 {
		moveCosts[i] = moveCosts[i-1] + i
	}

	for v := 0; v < maxPos; v++ {
		s1 := 0
		s2 := 0

		for _, p := range positions {
			s1 += common.Abs(p - v)
			s2 += moveCosts[common.Abs(p-v)]
		}

		if s2 < bestScore2 {
			bestScore2 = s2
		}
		if s1 < bestScore1 {
			bestScore1 = s1
		}

	}
	fmt.Println("Crabscore 1", bestScore1)
	fmt.Println("Crabscore 2", bestScore2)
}
