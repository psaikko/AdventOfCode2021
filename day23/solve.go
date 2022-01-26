package day23

import (
	"adventofcode/common"
	"fmt"
	"os"
)

func search(b burrow) (burrow, int) {

	heap := make(burrowHeap, 0)
	heap.insert(heapNode{state: b, cost: 0})
	seen := common.MakeIntSet()

	for {
		curr := heap.pop()
		h := curr.state.hash()
		if seen.Contains(h) {
			continue
		}
		seen.Put(h)

		if curr.state.atGoalState() {
			return curr.state, curr.cost
		}

		moves := curr.state.getMoves()
		for _, move := range moves {
			nextBurrow := curr.state.copy()
			nextBurrow.do(move)

			h := nextBurrow.hash()
			if !seen.Contains(h) {
				heap.insert(heapNode{
					state: nextBurrow,
					cost:  curr.cost + move.cost(),
				})
			}
		}
	}
}

func Run() {
	lines := common.ReadStringLines(os.Stdin)
	startState := make(burrow, 0)

	for _, line := range lines {
		startState = append(startState, []rune(line))
	}

	fmt.Println(startState)
	goalState, moveCost := search(startState)
	fmt.Println(goalState)
	fmt.Println(moveCost)
}
