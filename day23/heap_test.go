package day23

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestHeap(t *testing.T) {
	costs := make([]int, 100)
	for i := range costs {
		costs[i] = i
	}

	dummy := make(burrow, 0)
	for run := 0; run < 10; run++ {
		t.Run(fmt.Sprintf("run #%d", run), func(t2 *testing.T) {
			bh := make(burrowHeap, 0)

			rand.Shuffle(len(costs), func(i, j int) { costs[i], costs[j] = costs[j], costs[i] })
			for _, cost := range costs {
				bh.insert(heapNode{state: dummy, cost: cost})
			}

			// popped items should be in order
			for i := 0; i < len(costs); i++ {
				t2.Log(bh)
				j := bh.pop()
				if j.cost != i {
					t2.Log("expected", i, "was", j.cost)
					t2.Fail()
				}
			}
		})
	}
}
