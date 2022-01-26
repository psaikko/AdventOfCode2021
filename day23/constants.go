package day23

import "adventofcode/common"

const emptyTile = '.'
const wallTile = '#'

var amphipodTiles = common.MakeCharSet("ABCD")

const (
	hallwayY = 1
	roomTopY = 2
)

var roomX = map[rune]int{
	'A': 3,
	'B': 5,
	'C': 7,
	'D': 9,
}

var energyCosts = map[rune]int{
	'A': 1,
	'B': 10,
	'C': 100,
	'D': 1000,
}

var hallXPositions = common.MakeIntSet(1, 2, 4, 6, 8, 10, 11)
