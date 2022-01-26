package day23

import (
	"adventofcode/common"
	"fmt"
	"strings"
)

type burrow [][]rune

func (b burrow) String() string {
	builder := strings.Builder{}
	for y, line := range b {
		for x := range line {
			builder.WriteString(fmt.Sprintf("%c", b[y][x]))
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

func (b burrow) atGoalPosition(y, x int) bool {
	checkTile := b[y][x]
	if !amphipodTiles.Contains(checkTile) {
		panic(fmt.Errorf("checking %c at x=%d y=%d", checkTile, y, x))
	}
	return y >= roomTopY && x == roomX[checkTile]
}

func (b burrow) atGoalState() bool {
	for y, line := range b {
		for x, r := range line {
			if amphipodTiles.Contains(r) && !b.atGoalPosition(y, x) {
				return false
			}
		}
	}
	return true
}

type move struct {
	amphipodType rune
	from         common.IntPair
	to           common.IntPair
	dist         int
}

func (m move) reverse() move {
	return move{
		amphipodType: m.amphipodType,
		from:         m.to,
		to:           m.from,
		dist:         -m.dist,
	}
}

func (m move) cost() int {
	return energyCosts[m.amphipodType] * m.dist
}

func (b burrow) isValidFinalMove(m move) bool {
	// when moving into a room, make sure it's the correct room and position
	if m.to.X != roomX[m.amphipodType] {
		return false
	}
	tileBelow := b[m.to.Y+1][m.to.X]
	return tileBelow == wallTile || tileBelow == m.amphipodType
}

func (b burrow) isValid(m move) bool {
	// don't undo a valid final move
	if b.isValidFinalMove(m.reverse()) {
		return false
	}

	if m.from.Y >= roomTopY && m.to.Y == hallwayY {
		// moving from initial room to hallway
		// only stop in valid hallway positions
		return hallXPositions.Contains(m.to.X)
	} else {
		return b.isValidFinalMove(m)
	}
}

func (g burrow) getMovesFrom(y, x int) []move {
	amphipodType := g[y][x]
	moves := []move{}
	seen := make(common.IntSet, 0)
	posQueue := make([]common.IntPair, 0, 16)
	depthQueue := make([]int, 0, 16)

	startPos := common.MakeIntPair(x, y)
	seen.Put(startPos.Hash())
	posQueue = append(posQueue, startPos)
	depthQueue = append(depthQueue, 0)

	for len(posQueue) > 0 {
		p := posQueue[0]
		d := depthQueue[0]
		posQueue = posQueue[1:]
		depthQueue = depthQueue[1:]

		for _, next := range p.ManhattanNeighbors() {
			if g[next.Y][next.X] == emptyTile && !seen.Contains(next.Hash()) {
				seen.Put(next.Hash())
				posQueue = append(posQueue, next)
				depthQueue = append(depthQueue, d+1)
				nextMove := move{
					amphipodType: amphipodType,
					from:         common.MakeIntPair(x, y),
					to:           next,
					dist:         d + 1,
				}
				if g.isValid(nextMove) {
					moves = append(moves, nextMove)
				}
			}
		}
	}

	return moves
}

func (g burrow) getMoves() []move {
	moves := make([]move, 0)

	for y, line := range g {
		for x, tile := range line {
			if amphipodTiles.Contains(tile) {
				moves = append(moves, g.getMovesFrom(y, x)...)
			}
		}
	}

	return moves
}

func (g burrow) copy() burrow {
	newBurrow := make(burrow, len(g))
	for i, line := range g {
		newLine := make([]rune, len(line))
		copy(newLine, line)
		newBurrow[i] = newLine
	}
	return newBurrow
}

func (g burrow) do(m move) {
	if g[m.to.Y][m.to.X] != emptyTile {
		panic(fmt.Errorf("invalid move %+v", m))
	}
	if g[m.from.Y][m.from.X] != m.amphipodType {
		panic(fmt.Errorf("invalid move %+v", m))
	}

	g[m.to.Y][m.to.X] = m.amphipodType
	g[m.from.Y][m.from.X] = emptyTile
}

// hash implements djb2 for the string representation of the burrow
// collisions are possible but unlikely enough
func (g burrow) hash() int {
	hash := 5381
	for _, line := range g {
		for _, r := range line {
			hash = hash*33 + int(r)
		}
	}
	return hash
}
