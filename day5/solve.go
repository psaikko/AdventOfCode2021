package day5

import (
	"adventofcode/common"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type intPair struct {
	x int
	y int
}

func (from intPair) directionTo(to intPair) intPair {
	xd := to.x - from.x
	if xd != 0 {
		xd /= common.Abs(xd)
	}

	yd := to.y - from.y
	if yd != 0 {
		yd /= common.Abs(yd)
	}

	return intPair{
		x: xd,
		y: yd,
	}
}

func parseIntPair(s string) intPair {

	tokens := strings.Split(s, ",")
	ints := [2]int{}

	for i := range tokens {
		v, err := strconv.Atoi(tokens[i])
		if err != nil {
			panic(err)
		}
		ints[i] = v
	}

	return intPair{
		x: ints[0],
		y: ints[1],
	}
}

func Run() {
	const N = 1000
	ocean := [N][N]int{}

	lines := common.ReadStringsLines(os.Stdin)

	for _, line := range lines {
		to := parseIntPair(line[0])
		from := parseIntPair(line[2])

		d := from.directionTo(to)

		//if xd == 0 || yd == 0 {
		for {
			ocean[from.y][from.x] += 1

			if from.x == to.x && from.y == to.y {
				break
			}

			from.y += d.y
			from.x += d.x
		}
		//}
	}

	ct := 0
	for _, row := range ocean {
		for c := range row {
			if row[c] > 1 {
				ct++
			}
		}
	}

	fmt.Println("Overlapping lines:", ct)
}
