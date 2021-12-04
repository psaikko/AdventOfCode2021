package day1

import (
	"adventofcode/common"
	"fmt"
	"os"
)

// Run reads input from Stdin to solve https://adventofcode.com/2021/day/1
func Run() {
	vals := common.ReadIntLines(os.Stdin)

	count := 0
	for i, v := range vals {
		if i > 0 && v > vals[i-1] {
			count += 1
		}
	}

	fmt.Printf("Depth increases %d times\n", count)

	count = 0
	sum := 0

	for i, v := range vals {
		if i < 3 {
			sum += v
		} else {
			newSum := sum - vals[i-3] + v
			if newSum > sum {
				count += 1
			}
			sum = newSum
		}
	}

	fmt.Printf("Window depth increases %d times\n", count)
}
