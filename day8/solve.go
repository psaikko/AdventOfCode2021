package day8

import (
	"adventofcode/common"
	"fmt"
	"os"
	"strings"
)

func commonChars(a, b string) int {
	ct := 0
	for _, c := range a {
		for _, c2 := range b {
			if c == c2 {
				ct++
			}
		}
	}
	return ct
}

func charSetEqual(a, b string) bool {
	return len(a) == len(b) && commonChars(a, b) == len(a)
}

func Run() {
	lines := common.ReadStringLines(os.Stdin)
	totalOut := 0
	count1478 := 0
	for _, line := range lines {
		parts := strings.Split(strings.TrimSpace(line), " | ")
		inValues := strings.Split(parts[0], " ")
		outValues := strings.Split(parts[1], " ")
		codes := [10]string{}

		// easy cases
		for _, v := range inValues {
			switch len(v) {
			case 2:
				codes[1] = v
			case 7:
				codes[8] = v
			case 3:
				codes[7] = v
			case 4:
				codes[4] = v
			}
		}

		// tricky cases
		for _, v := range inValues {
			switch len(v) {
			case 5:
				if commonChars(v, codes[7]) == 3 {
					codes[3] = v
				} else if commonChars(v, codes[4]) == 2 {
					codes[2] = v
				} else {
					codes[5] = v
				}
			case 6:
				if commonChars(v, codes[7]) == 2 {
					codes[6] = v
				} else if commonChars(v, codes[4]) == 4 {
					codes[9] = v
				} else {
					codes[0] = v
				}
			}
		}

		// do counts
		val := 0
		m := [4]int{1000, 100, 10, 1}
		for i, v1 := range outValues {
			for k, v2 := range codes {
				if charSetEqual(v1, v2) {
					val += m[i] * k
					if k == 1 || k == 4 || k == 7 || k == 8 {
						count1478++
					}
					break
				}
			}

		}
		totalOut += val
	}

	fmt.Println(count1478)
	fmt.Println(totalOut)
}
