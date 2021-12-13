package day12

import (
	"adventofcode/common"
	"fmt"
	"os"
	"strings"
)

const startNode = "start"
const endNode = "end"

func isSmall(node string) bool {
	return node[0] >= 'a' && node[0] <= 'z'
}

func countPathsFrom(from string, visits map[string]int, edges map[string][]string, flag bool) int {

	if from == endNode {
		return 1
	}

	count := 0
	for _, next := range edges[from] {

		if flag && visits[next] == 1 {
			continue
		}
		if visits[next] == 2 {
			continue
		}

		nextFlag := flag
		if isSmall(next) {
			visits[next]++
			if visits[next] == 2 {
				nextFlag = true
			}
		}

		count += countPathsFrom(next, visits, edges, nextFlag)

		if isSmall(next) {
			visits[next]--
		}
	}

	return count
}

func countPaths(edges map[string][]string, canVisitTwice bool) int {
	visits := make(map[string]int)
	if canVisitTwice {
		visits[startNode] = 2
	} else {
		visits[startNode] = 1
	}
	return countPathsFrom(startNode, visits, edges, !canVisitTwice)
}

func Run() {
	lines := common.ReadStringLines(os.Stdin)

	edges := make(map[string][]string)
	nodes := common.MakeStringSet()

	for _, line := range lines {
		tokens := strings.Split(line, "-")
		from := tokens[0]
		to := tokens[1]

		edges[to] = append(edges[to], from)
		edges[from] = append(edges[from], to)
		nodes.Put(from)
		nodes.Put(to)
	}

	fmt.Println("Small caves once:", countPaths(edges, false))
	fmt.Println("Small cave twice:", countPaths(edges, true))
}
