package main

import (
	"adventofcode/day1"
	"adventofcode/day2"
	"adventofcode/day3"
	"adventofcode/day4"
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
	case "2":
		day2.Run()
	case "3":
		day3.Run()
	case "4":
		day4.Run()
	default:
		panic(fmt.Sprintf("Day '%s' not implemented", day))
	}
}
