package main

import (
	"adventofcode/day1"
	"adventofcode/day10"
	"adventofcode/day11"
	"adventofcode/day12"
	"adventofcode/day13"
	"adventofcode/day14"
	"adventofcode/day15"
	"adventofcode/day16"
	"adventofcode/day17"
	"adventofcode/day18"
	"adventofcode/day2"
	"adventofcode/day20"
	"adventofcode/day21"
	"adventofcode/day3"
	"adventofcode/day4"
	"adventofcode/day5"
	"adventofcode/day6"
	"adventofcode/day7"
	"adventofcode/day8"
	"adventofcode/day9"
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
	case "5":
		day5.Run()
	case "6":
		day6.Run()
	case "7":
		day7.Run()
	case "8":
		day8.Run()
	case "9":
		day9.Run()
	case "10":
		day10.Run()
	case "11":
		day11.Run()
	case "12":
		day12.Run()
	case "13":
		day13.Run()
	case "14":
		day14.Run()
	case "15":
		day15.Run()
	case "16":
		day16.Run()
	case "17":
		day17.Run()
	case "18":
		day18.Run()
	case "20":
		day20.Run()
	case "21":
		day21.Run()
	default:
		panic(fmt.Sprintf("Day '%s' not implemented", day))
	}
}
