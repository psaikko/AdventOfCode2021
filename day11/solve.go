package day11

import (
	"adventofcode/common"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func neighborPositions(mat [][]int, y int, x int) [][2]int {
	n := make([][2]int, 0)

	if x > 0 {
		n = append(n, [2]int{y, x - 1})
		if y > 0 {
			n = append(n, [2]int{y - 1, x - 1})
		}
		if y < len(mat)-1 {
			n = append(n, [2]int{y + 1, x - 1})
		}
	}
	if x < len(mat[0])-1 {
		n = append(n, [2]int{y, x + 1})
		if y > 0 {
			n = append(n, [2]int{y - 1, x + 1})
		}
		if y < len(mat)-1 {
			n = append(n, [2]int{y + 1, x + 1})
		}
	}
	if y > 0 {
		n = append(n, [2]int{y - 1, x})
	}
	if y < len(mat)-1 {
		n = append(n, [2]int{y + 1, x})
	}

	return n
}

func flash(mat [][]int, y int, x int) int {
	flashes := 0

	if mat[y][x] > 9 {
		flashes++
		mat[y][x] = 0
		for _, n := range neighborPositions(mat, y, x) {
			y2 := n[0]
			x2 := n[1]
			if mat[y2][x2] != 0 {
				mat[y2][x2]++
				flashes += flash(mat, y2, x2)
			}
		}
	}

	return flashes
}

func Run() {
	lines := common.ReadStringLines(os.Stdin)

	energies := make([][]int, len(lines))
	for k, line := range lines {
		s := strings.Split(line, "")
		is := make([]int, len(s))
		for j, c := range s {
			i, err := strconv.Atoi(c)
			if err != nil {
				panic(err)
			}
			is[j] = i
		}
		energies[k] = is
	}

	flashes := 0

	for i := 1; ; i++ {
		for y := 0; y < 10; y++ {
			for x := 0; x < 10; x++ {
				energies[y][x]++
			}
		}

		stepFlashes := 0
		for y := 0; y < 10; y++ {
			for x := 0; x < 10; x++ {
				stepFlashes += flash(energies, y, x)
			}
		}
		flashes += stepFlashes

		if stepFlashes == 100 {
			fmt.Println("Synchronized at", i)
			break
		}

		if i == 100 {
			fmt.Println("After 100 steps", flashes)
		}
	}
}
