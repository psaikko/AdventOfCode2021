package day9

import (
	"adventofcode/common"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func neighborValues(mat [][]int, ip common.IntPair) []int {
	n := make([]int, 0)
	x, y := ip.X, ip.Y

	if x > 0 {
		n = append(n, mat[y][x-1])
	}
	if x < len(mat[0])-1 {
		n = append(n, mat[y][x+1])
	}
	if y > 0 {
		n = append(n, mat[y-1][x])
	}
	if y < len(mat)-1 {
		n = append(n, mat[y+1][x])
	}

	return n
}

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

func Run() {
	lines := common.ReadStringLines(os.Stdin)

	heightMap := make([][]int, len(lines))
	for k, line := range lines {
		s := strings.Split(line, "")
		row := make([]int, len(s))
		for j, c := range s {
			i, err := strconv.Atoi(c)
			if err != nil {
				panic(err)
			}
			row[j] = i
		}
		heightMap[k] = row
	}

	riskSum := 0
	lowPoints := []common.IntPair{}

	for y := range heightMap {
		for x := range heightMap[y] {
			h := heightMap[y][x]
			point := common.MakeIntPair(x, y)
			if h < common.Min(neighborValues(heightMap, point)...) {
				riskSum += h + 1
				lowPoints = append(lowPoints, point)
			}
		}
	}

	fmt.Println("Total risk level", riskSum)
	seen := make(common.IntSet)
	basinSizes := []int{}

	for _, lp := range lowPoints {
		size := 0
		if seen.Contains(lp.Hash()) {
			continue // part of another basin
		}
		q := []common.IntPair{lp}

		for len(q) > 0 {
			next := q[0]
			q = q[1:]

			if !seen.Contains(next.Hash()) {
				if heightMap[next.Y][next.X] < 9 {
					q = append(q, neighborPositions(heightMap, next)...)
					seen.Put(next.Hash())
					size += 1
				}
			}
		}

		basinSizes = append(basinSizes, size)
	}

	sort.Ints(basinSizes)
	n := len(basinSizes)
	fmt.Println("Largest basins prod", basinSizes[n-1]*basinSizes[n-2]*basinSizes[n-3])
}
