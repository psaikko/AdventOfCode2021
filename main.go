package main

import (
	"adventofcode/day1"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("Need day as first arg")
	}
	switch day := os.Args[1]; day {
	case "1":
		day1.Run()
	default:
		panic(fmt.Sprintf("Day '%s' not implemented", day))
	}
}
