package day8

import (
	"adventofcode/common"
	"fmt"
	"os"
	"strings"
)

func Run() {
	lines := common.ReadStringLines(os.Stdin)
	totalOut := 0
	count1478 := 0
	for _, line := range lines {
		parts := strings.Split(strings.TrimSpace(line), " | ")
		inValues := strings.Split(parts[0], " ")
		outValues := strings.Split(parts[1], " ")
		sets := [10]common.CharSet{}

		// easy cases
		for _, v := range inValues {
			cs := common.MakeCharSet(v)
			switch len(cs) {
			case 2:
				sets[1] = cs
			case 7:
				sets[8] = cs
			case 3:
				sets[7] = cs
			case 4:
				sets[4] = cs
			}
		}

		// tricky cases
		for _, v := range inValues {
			cs := common.MakeCharSet(v)
			switch len(cs) {
			case 5:
				if len(cs.Intersection(sets[7])) == 3 {
					sets[3] = cs
				} else if len(cs.Intersection(sets[4])) == 2 {
					sets[2] = cs
				} else {
					sets[5] = cs
				}
			case 6:
				if len(cs.Intersection(sets[7])) == 2 {
					sets[6] = cs
				} else if len(cs.Intersection(sets[4])) == 4 {
					sets[9] = cs
				} else {
					sets[0] = cs
				}
			}
		}

		// do counts
		val := 0
		m := [4]int{1000, 100, 10, 1}
		for i, v1 := range outValues {
			cs := common.MakeCharSet(v1)
			for k, kSet := range sets {
				if cs.Equals(kSet) {
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
