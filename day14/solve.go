package day14

import (
	"adventofcode/common"
	"fmt"
	"os"
	"strings"
)

func elementScore(steps int, template string, rules map[string]string) int {
	pairCounts := make(map[string]int)

	for j := 0; j < len(template)-1; j++ {
		pair := template[j : j+2]
		pairCounts[pair] = 1
	}

	for i := 0; i < steps; i++ {
		newPairCounts := make(map[string]int, 0)

		for p := range pairCounts {
			p1 := string(p[0]) + rules[p]
			p2 := rules[p] + string(p[1])
			newPairCounts[p1] += pairCounts[p]
			newPairCounts[p2] += pairCounts[p]
		}

		pairCounts = newPairCounts
	}

	count := make(map[rune]int)

	for pair := range pairCounts {
		ct := pairCounts[pair]
		for _, r := range pair {
			count[r] += ct
		}
	}

	count[rune(template[0])]++
	count[rune(template[len(template)-1])]++

	least := -1
	most := 0
	for r := range count {
		if count[r] < least || least == -1 {
			least = count[r]
		}
		if count[r] > most {
			most = count[r]
		}
	}

	return (most - least) / 2
}

func Run() {
	lines := common.ReadStringLines(os.Stdin)
	template := lines[0]
	rulesLines := lines[2:]

	rules := make(map[string]string)

	for _, line := range rulesLines {
		parts := strings.Split(line, " -> ")
		rules[parts[0]] = parts[1]
	}

	fmt.Println("After 10 steps", elementScore(10, template, rules))
	fmt.Println("After 40 steps", elementScore(40, template, rules))
}
