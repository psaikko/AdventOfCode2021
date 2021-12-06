package day6

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func fishSum(a [10]int) int {
	s := 0
	for _, v := range a {
		s += v
	}
	return s
}

func Run() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	numbersLine := scanner.Text()
	timerCounts := [10]int{}
	for _, s := range strings.Split(numbersLine, ",") {
		v, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		timerCounts[v]++
	}

	for i := 0; i < 256; i++ {
		for j := 0; j < len(timerCounts); j++ {
			v := timerCounts[j]
			if j == 0 {
				timerCounts[9] += v
				timerCounts[7] += v
			} else {
				timerCounts[j-1] += v
			}
			timerCounts[j] = 0
		}
		if i == 80 {
			fmt.Printf("After %d days, %d fish\n", 80, fishSum(timerCounts))
		}
	}

	fmt.Printf("After %d days, %d fish\n", 256, fishSum(timerCounts))
}
