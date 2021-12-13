package day5

import (
	"adventofcode/common"
	"fmt"
	"os"
)

func direction(from, to common.IntPair) common.IntPair {
	xd := to.X - from.X
	if xd != 0 {
		xd /= common.Abs(xd)
	}

	yd := to.Y - from.Y
	if yd != 0 {
		yd /= common.Abs(yd)
	}

	return common.MakeIntPair(xd, yd)
}

func Run() {
	const N = 1000
	ocean := [N][N]int{}

	lines := common.ReadStringsLines(os.Stdin)

	for _, line := range lines {
		to := common.ParseIntPair(line[0])
		from := common.ParseIntPair(line[2])

		d := direction(from, to)

		//if xd == 0 || yd == 0 {
		for {
			ocean[from.Y][from.X] += 1

			if from.X == to.X && from.Y == to.Y {
				break
			}

			from.Y += d.Y
			from.X += d.X
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
