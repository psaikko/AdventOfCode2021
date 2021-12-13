package common

import (
	"strconv"
	"strings"
)

type IntPair struct {
	X int
	Y int
}

func MakeIntPair(x, y int) IntPair {
	return IntPair{X: x, Y: y}
}

func ParseIntPair(s string) IntPair {
	parts := strings.Split(s, ",")
	if len(parts) != 2 {
		panic("Expected two comma separated values")
	}

	x, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}

	y, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	return IntPair{X: x, Y: y}
}
