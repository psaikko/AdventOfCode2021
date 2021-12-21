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

func (ip IntPair) Hash() int {
	return ip.X*1000000 + ip.Y
}

type IntTriple struct {
	X int
	Y int
	Z int
}

func MakeIntTriple(x, y, z int) IntTriple {
	return IntTriple{X: x, Y: y, Z: z}
}

func ParseIntTriple(s string) IntTriple {
	parts := strings.Split(s, ",")
	if len(parts) != 3 {
		panic("Expected three comma separated values")
	}

	x, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}

	y, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	z, err := strconv.Atoi(parts[2])
	if err != nil {
		panic(err)
	}

	return IntTriple{X: x, Y: y, Z: z}
}

func (it IntTriple) Hash() int {
	return it.X*10000*10000 + it.Y*10000 + it.Z
}

func (it IntTriple) DistanceTo(it2 IntTriple) int {
	return Abs(it.X-it2.X) + Abs(it.Y-it2.Y) + Abs(it.Z-it2.Z)
}
