package day19

import (
	"adventofcode/common"
	"fmt"
	"os"
)

type scanner []common.IntTriple

func matches(s1, s2 scanner) int {
	ints := make(common.IntSet, 0)
	for _, it := range s1 {
		ints.Put(it.Hash())
	}
	count := 0
	for _, it := range s2 {
		if ints.Contains(it.Hash()) {
			count++
		}
	}
	return count
}

func (s scanner) apply(t transform) scanner {
	return s.rot(t.xr, t.yr, t.zr).shift(t.xs, t.ys, t.zs)
}

func (s scanner) shift(x, y, z int) scanner {
	newScanner := make(scanner, len(s))
	for i := 0; i < len(s); i++ {
		newScanner[i] = common.MakeIntTriple(
			s[i].X+x,
			s[i].Y+y,
			s[i].Z+z,
		)
	}
	return newScanner
}

func (s scanner) rot(x, y, z int) scanner {
	newScanner := make(scanner, len(s))
	for i := 0; i < len(s); i++ {
		newScanner[i] = s[i]
	}
	for i := 0; i < x; i++ {
		for j := 0; j < len(s); j++ {
			newScanner[j] = rotX(newScanner[j])
		}
	}
	for i := 0; i < y; i++ {
		for j := 0; j < len(s); j++ {
			newScanner[j] = rotY(newScanner[j])
		}
	}
	for i := 0; i < z; i++ {
		for j := 0; j < len(s); j++ {
			newScanner[j] = rotZ(newScanner[j])
		}
	}
	return newScanner
}

func rotX(it common.IntTriple) common.IntTriple {
	return common.MakeIntTriple(it.X, it.Z, -it.Y)
}

func rotY(it common.IntTriple) common.IntTriple {
	return common.MakeIntTriple(it.Z, it.Y, -it.X)
}

func rotZ(it common.IntTriple) common.IntTriple {
	return common.MakeIntTriple(-it.Y, it.X, it.Z)
}

type transform struct {
	xr, yr, zr, xs, ys, zs int
}

func fit(s1, s2 scanner) *transform {

	for xr := 0; xr < 4; xr++ {
		for yr := 0; yr < 4; yr++ {
			for zr := 0; zr < 4; zr++ {
				rotatedScanner := s2.rot(xr, yr, zr)

				for _, b2 := range rotatedScanner {
					for _, b1 := range s1 {
						shiftedRotatedScanner := rotatedScanner.shift(
							b1.X-b2.X,
							b1.Y-b2.Y,
							b1.Z-b2.Z)
						ms := matches(s1, shiftedRotatedScanner)
						if ms == 12 {
							return &transform{
								xr, yr, zr,
								b1.X - b2.X,
								b1.Y - b2.Y,
								b1.Z - b2.Z,
							}
						}
					}

				}
			}
		}
	}

	return nil
}

func Run() {
	lines := common.ReadStringLines(os.Stdin)

	scanners := []scanner{}
	tail := -1

	newLine := true
	for _, line := range lines {
		if newLine {
			newLine = false
			scanners = append(scanners, make(scanner, 0))
			tail = len(scanners) - 1
			continue
		}
		if line == "" {
			newLine = true
			continue
		}

		fmt.Println(line)
		beacon := common.ParseIntTriple(line)
		scanners[tail] = append(scanners[tail], beacon)
	}

	fmt.Println(scanners)

	fmt.Println(scanners[0].rot(1, 0, 0))

	fmt.Println(scanners[0].shift(1, 0, 0))

	isFixed := make([]bool, len(scanners))
	isFixed[0] = true
	nFixed := 1

	scannerPositions := []common.IntTriple{common.MakeIntTriple(0, 0, 0)}

	for nFixed < len(scanners) {
	outer:
		for i := 0; i < len(scanners); i++ {
			if isFixed[i] {
				for j := 0; j < len(scanners); j++ {
					if !isFixed[j] {
						t := fit(scanners[i], scanners[j])
						if t != nil {
							fmt.Println(i, j, *t)
							scanners[j] = scanners[j].apply(*t)
							isFixed[j] = true
							nFixed++
							scannerPositions = append(scannerPositions,
								common.MakeIntTriple(t.xs, t.ys, t.zs),
							)
							break outer
						}
					}
				}
			}
		}
	}

	beaconHashes := make(common.IntSet)

	for _, sc := range scanners {
		for _, bc := range sc {
			beaconHashes.Put(bc.Hash())
		}
	}

	fmt.Println(len(beaconHashes))

	maxDist := 0

	for i := range scannerPositions {
		for j := i + 1; j < len(scannerPositions); j++ {
			dist := scannerPositions[i].DistanceTo(scannerPositions[j])
			if dist > maxDist {
				fmt.Println(scannerPositions[i], scannerPositions[j])
				maxDist = dist
			}
		}
	}

	fmt.Println(maxDist)
}
