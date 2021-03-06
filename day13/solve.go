package day13

import (
	"adventofcode/common"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	lines := common.ReadStringLines(os.Stdin)

	dots := []common.IntPair{}
	dotsMap := common.IntSet{}

	foldsStart := 0
	for i, line := range lines {
		if line == "" {
			foldsStart = i
			break
		}
		dot := common.ParseIntPair(line)
		dots = append(dots, dot)
		dotsMap.Put(dot.Hash())
	}

	for i, fold := range lines[foldsStart+1:] {
		toks := strings.Split(strings.TrimPrefix(fold, "fold along "), "=")
		kind := toks[0]
		val, _ := strconv.Atoi(toks[1])

		switch kind {
		case "x":
			for i := range dots {
				dot := dots[i]
				if dot.X > val {
					delete(dotsMap, dot.Hash())
					dots[i].X = (val - (dot.X - val))
					dotsMap.Put(dots[i].Hash())
				}
			}
		case "y":
			for i := range dots {
				dot := dots[i]
				if dot.Y > val {
					delete(dotsMap, dot.Hash())
					dots[i].Y = (val - (dot.Y - val))
					dotsMap.Put(dots[i].Hash())
				}
			}
		}
		if i == 0 {
			fmt.Println(len(dotsMap))
		}
	}

	img := [100][100]int{}

	for _, dot := range dots {
		img[dot.Y][dot.X] = 1
	}

	for _, line := range img[:6] {
		for _, chr := range line[:40] {
			if chr == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
