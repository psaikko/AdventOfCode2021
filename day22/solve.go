package day22

import (
	"adventofcode/common"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const minCoord = -50
const maxCoord = 50

type cuboid struct {
	on         bool
	xMin, xMax int
	yMin, yMax int
	zMin, zMax int
}

func parseLimits(s string) (min, max int) {
	lims := strings.Split(s[2:], "..")
	min, _ = strconv.Atoi(lims[0])
	max, _ = strconv.Atoi(lims[1])
	return
}

func hash(x, y, z int) int {
	return 1000*1000*x + 1000*y + z
}

func Run() {
	lines := common.ReadStringLines(os.Stdin)

	cuboids := []cuboid{}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		on := false
		if parts[0] == "on" {
			on = true
		}
		coordParts := strings.Split(parts[1], ",")
		xMin, xMax := parseLimits(coordParts[0])
		yMin, yMax := parseLimits(coordParts[1])
		zMin, zMax := parseLimits(coordParts[2])

		cuboids = append(cuboids, cuboid{
			on, xMin, xMax, yMin, yMax, zMin, zMax,
		})
	}

	active := make(common.IntSet, 0)
	for _, cb := range cuboids {
		for x := common.Max(minCoord, cb.xMin); x <= common.Min(maxCoord, cb.xMax); x++ {
			for y := common.Max(minCoord, cb.yMin); y <= common.Min(maxCoord, cb.yMax); y++ {
				for z := common.Max(minCoord, cb.zMin); z <= common.Min(maxCoord, cb.zMax); z++ {
					if cb.on {
						active.Put(hash(x, y, z))
					} else {
						delete(active, hash(x, y, z))
					}
				}
			}
		}
	}
	fmt.Println(len(active))

}
