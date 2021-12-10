package day10

import (
	"adventofcode/common"
	"fmt"
	"os"
	"sort"
	"strings"
)

func Run() {

	scoreMap := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	completeMap := map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}

	closingPair := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
		"<": ">",
	}

	lines := common.ReadStringLines(os.Stdin)
	errorScore := 0
	completionScores := []int{}

outerLoop:
	for _, line := range lines {
		q := []string{}
		for _, r := range strings.Split(line, "") {
			switch r {
			case "(", "[", "{", "<":
				q = append(q, closingPair[r])
			case ")", "]", "}", ">":
				if q[len(q)-1] != r {
					errorScore += scoreMap[r]
					continue outerLoop
				} else {
					q = q[:len(q)-1]
				}
			}
		}
		lineScore := 0
		for i := len(q) - 1; i >= 0; i-- {
			lineScore = lineScore*5 + completeMap[q[i]]
		}
		completionScores = append(completionScores, lineScore)
	}

	fmt.Println("Error score", errorScore)
	sort.Ints(completionScores)
	fmt.Println("Middle score", completionScores[len(completionScores)/2])
}
