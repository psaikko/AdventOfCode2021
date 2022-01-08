package day17

import (
	"adventofcode/common"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parsePart(s string) (min, max int) {
	parts := strings.Split(s[2:], "..")
	min, errMin := strconv.Atoi(parts[0])
	if errMin != nil {
		panic(errMin)
	}
	max, errMax := strconv.Atoi(parts[1])
	if errMax != nil {
		panic(errMax)
	}
	return
}

func Run() {
	line := common.ReadStringLines(os.Stdin)[0]
	line = strings.TrimPrefix(line, "target area: ")
	parts := strings.Split(line, ", ")
	xMin, xMax := parsePart(parts[0])
	yMin, yMax := parsePart(parts[1])

	bestHighest := 0
	okVelocities := 0

	for yStartV := -1000; yStartV < 1000; yStartV++ {
	nextVal:
		for xStartV := 0; xStartV < 1000; xStartV++ {

			x, y := 0, 0
			xv, yv := xStartV, yStartV
			highest := 0

			for x <= xMax && y >= yMin {
				x += xv
				y += yv

				highest = common.Max(y, highest)

				yv -= 1
				if xv > 0 {
					xv -= 1
				}
				if xv < 0 {
					xv += 1
				}

				if x >= xMin && x <= xMax && y >= yMin && y <= yMax {
					bestHighest = common.Max(highest, bestHighest)
					okVelocities++
					continue nextVal
				}
			}
		}
	}

	fmt.Println("Highest y position", bestHighest)
	fmt.Println("Initial velocities", okVelocities)
}
