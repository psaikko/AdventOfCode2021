package day2

import (
	"adventofcode/common"
	"fmt"
	"os"
	"strconv"
)

func Run() {
	lines := common.ReadStringsLines(os.Stdin)

	x := 0
	y := 0
	aim := 0

	for _, tokens := range lines {
		amount, _ := strconv.Atoi(tokens[1])
		switch tokens[0] {
		case "forward":
			x += amount
			y += aim * amount
		case "backward":
			x -= amount
		case "down":
			aim += amount
		case "up":
			aim -= amount
		}
	}

	fmt.Println("Part 1", x*aim)
	fmt.Println("Part 2", x*y)
}
