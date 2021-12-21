package day20

import (
	"adventofcode/common"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func address(grid [][]byte, r, c int, pad byte) int {

	addr := 0

	fill := 0
	if pad == '#' {
		fill = 1
	}

	for ir := r - 1; ir <= r+1; ir++ {
		for ic := c - 1; ic <= c+1; ic++ {
			addr <<= 1

			if ir < 0 || ir > len(grid)-1 {
				addr += fill
				continue
			}

			if ic < 0 || ic > len(grid[0])-1 {
				addr += fill
				continue
			}

			if grid[ir][ic] == '#' {
				addr += 1
			}
		}
	}

	return addr
}

const p = 2

func pad(lines [][]byte, padc byte) [][]byte {
	n := len(lines[0])

	newLines := [][]byte{}
	for i := 0; i < p; i++ {
		newLines = append(newLines, []byte(strings.Repeat(string(padc), n+p*2)))
	}
	linePad := bytes.Repeat([]byte{'.'}, p)

	for _, line := range lines {
		padLine := append(linePad, append(line, linePad...)...)
		newLines = append(newLines, padLine)
	}

	for i := 0; i < p; i++ {
		newLines = append(newLines, []byte(strings.Repeat(string(padc), n+p*2)))
	}
	return newLines
}

func Run() {
	lines := common.ReadStringLines(os.Stdin)
	lookup := lines[0]
	grid := [][]byte{}
	for _, line := range lines[2:] {
		grid = append(grid, []byte(line))
	}

	var padChar byte = '.'

	for i := 0; i < 50; i++ {

		nextGrid := pad(grid, padChar)
		for r := -p; r < len(grid)+p; r++ {
			for c := -p; c < len(grid[0])+p; c++ {
				addr := address(grid, r, c, padChar)
				char := lookup[addr]
				nextGrid[r+p][c+p] = char
			}
		}

		grid = nextGrid
		if padChar == '.' {
			padChar = lookup[0b000000000]
		} else {
			padChar = lookup[0b111111111]
		}
	}

	ct := 0

	for _, line := range grid {
		for _, char := range line {
			if char == '#' {
				ct++
			}
		}
	}

	fmt.Println(ct)
}
