package day25

import (
	"adventofcode/common"
	"fmt"
	"os"
)

const (
	empty = iota
	right
	down
)

type grid [][]int

func (g grid) copyTo(g2 grid) {
	for r, row := range g {
		for c := range row {
			g2[r][c] = g[r][c]
		}
	}
}

func Run() {
	lines := common.ReadStringLines(os.Stdin)
	h := len(lines)
	w := len(lines[0])
	seafloor := make(grid, h)
	tmp := make(grid, h)
	for i := 0; i < h; i++ {
		seafloor[i] = make([]int, w)
		tmp[i] = make([]int, w)
		for j, c := range lines[i] {
			switch c {
			case '>':
				seafloor[i][j] = right
			case 'v':
				seafloor[i][j] = down
			case '.':
				seafloor[i][j] = empty
			default:
				panic(c)
			}
		}
	}
	seafloor.copyTo(tmp)

	moved := true
	var i int
	for i = 0; moved; i++ {
		moved = false
		// move cucumber right
		for r, row := range seafloor {
			for c, val := range row {
				if val == right {
					if seafloor[r][(c+1)%w] == empty {
						moved = true
						tmp[r][c] = empty
						tmp[r][(c+1)%w] = right
					}
				}
			}
		}
		tmp.copyTo(seafloor)

		// move cucumber down
		for c := 0; c < w; c++ {
			for r := 0; r < h; r++ {
				if seafloor[r][c] == down {
					if seafloor[(r+1)%h][c] == empty {
						moved = true
						tmp[(r+1)%h][c] = down
						tmp[r][c] = empty
					}
				}
			}
		}

		tmp.copyTo(seafloor)
	}

	fmt.Println("Stops moving after", i, "steps")
}
