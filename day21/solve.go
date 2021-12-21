package day21

import (
	"adventofcode/common"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	lines := common.ReadStringLines(os.Stdin)

	initPositions := [2]int{}
	positions := [2]int{}
	scores := [2]int{}

	p1Parts := strings.Split(lines[0], ": ")
	initPositions[0], _ = strconv.Atoi(p1Parts[1])

	p2Parts := strings.Split(lines[1], ": ")
	initPositions[1], _ = strconv.Atoi(p2Parts[1])

	positions = initPositions

	rolls := 0
	turn := 0

	for scores[0] < 1000 && scores[1] < 1000 {
		move := 0
		for i := 0; i < 3; i++ {
			rolls++
			move += ((rolls - 1) % 100) + 1
		}
		positions[turn] = ((positions[turn] + move - 1) % 10) + 1
		scores[turn] += positions[turn]
		turn = (turn + 1) % 2
	}
	fmt.Println("Part 1", common.Min(scores[:]...)*rolls)

	moveDistribution := [10]int{}
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				moveDistribution[i+j+k]++
			}
		}
	}

	const minMove = 3
	const maxMove = 9
	const maxScore = 21

	cache := [11][11][maxScore][maxScore][2][2]int{}
	isCached := [11][11][maxScore][maxScore][2]bool{}

	var winsFunc func(int, int, int, int, int) [2]int
	winsFunc = func(p1pos, p2pos, p1score, p2score, turn int) [2]int {
		if p1score >= maxScore {
			return [2]int{1, 0}
		}

		if p2score >= maxScore {
			return [2]int{0, 1}
		}

		if isCached[p1pos][p2pos][p1score][p2score][turn] {
			return cache[p1pos][p2pos][p1score][p2score][turn]
		}

		totalWins := [2]int{}
		for move := minMove; move <= maxMove; move++ {
			moveCount := moveDistribution[move]
			if turn == 0 {
				nextPos := ((p1pos + move - 1) % 10) + 1
				nextScore := p1score + nextPos
				wins := winsFunc(nextPos, p2pos, nextScore, p2score, 1)
				totalWins[0] += wins[0] * moveCount
				totalWins[1] += wins[1] * moveCount
			} else {
				nextPos := ((p2pos + move - 1) % 10) + 1
				nextScore := p2score + nextPos
				wins := winsFunc(p1pos, nextPos, p1score, nextScore, 0)
				totalWins[0] += wins[0] * moveCount
				totalWins[1] += wins[1] * moveCount
			}
		}

		cache[p1pos][p2pos][p1score][p2score][turn] = totalWins
		isCached[p1pos][p2pos][p1score][p2score][turn] = true
		return totalWins
	}

	totalWins := winsFunc(initPositions[0], initPositions[1], 0, 0, 0)
	fmt.Println("Part 2", common.Max(totalWins[:]...))
}
