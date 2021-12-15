package day15

import (
	"adventofcode/common"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func neighborPositions(mat [][]int, ip common.IntPair) []common.IntPair {
	x, y := ip.X, ip.Y
	n := []common.IntPair{}

	if x > 0 {
		n = append(n, common.MakeIntPair(x-1, y))
	}
	if x < len(mat[0])-1 {
		n = append(n, common.MakeIntPair(x+1, y))
	}
	if y > 0 {
		n = append(n, common.MakeIntPair(x, y-1))
	}
	if y < len(mat)-1 {
		n = append(n, common.MakeIntPair(x, y+1))
	}

	return n
}

func incr(row []int, k int) []int {
	newRow := make([]int, len(row))
	for i := range row {
		newRow[i] = row[i] + k
		if newRow[i] > 9 {
			newRow[i] -= 9
		}
	}
	return newRow
}

func dupRow(row []int) []int {
	n := len(row)

	for i := 0; i < 4; i++ {
		row = append(row, incr(row[:n], i+1)...)
	}

	return row
}

func enlarge(arr [][]int) [][]int {
	m := len(arr)

	for i := 0; i < m; i++ {
		arr[i] = dupRow(arr[i])
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < m; j++ {
			arr = append(arr, incr(arr[j], i+1))
		}
	}

	return arr
}

func shortestPath(from, to common.IntPair, mapGrid [][]int) int {
	q := []common.IntPair{from}
	c := make(map[common.IntPair]int)
	c[from] = 0

	for len(q) > 0 {
		curr := q[0]
		currDist := c[curr]
		q = q[1:]

		for _, next := range neighborPositions(mapGrid, curr) {
			nextDanger := mapGrid[next.Y][next.X]
			dist, seen := c[next]
			if seen {
				if currDist+nextDanger < dist {
					q = append(q, next)
					c[next] = currDist + nextDanger
				}
			} else {
				q = append(q, next)
				c[next] = currDist + nextDanger
			}
		}
	}

	return c[to]
}

func Run() {
	lines := common.ReadStringLines(os.Stdin)
	danger := make([][]int, len(lines))
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
		danger[k] = is
	}

	from := common.MakeIntPair(0, 0)
	to := common.MakeIntPair(len(danger[0])-1, len(danger)-1)

	fmt.Println(shortestPath(from, to, danger))

	danger = enlarge(danger)
	to2 := common.MakeIntPair(len(danger[0])-1, len(danger)-1)
	fmt.Println(shortestPath(from, to2, danger))
}
