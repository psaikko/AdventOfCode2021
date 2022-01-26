package day23

import (
	"fmt"
	"strings"
)

type heapNode struct {
	state burrow
	cost  int
}

type burrowHeap []heapNode

func heapLeftChild(i int) int {
	return i*2 + 1
}

func heapRightChild(i int) int {
	return i*2 + 2
}

func heapParent(i int) int {
	return (i - 1) / 2
}

func (bh burrowHeap) String() string {
	b := strings.Builder{}
	for _, gc := range bh {
		b.WriteString(fmt.Sprintf("%d ", gc.cost))
	}
	return b.String()
}

func (bh *burrowHeap) swap(i, j int) {
	(*bh)[j], (*bh)[i] = (*bh)[i], (*bh)[j]
}

func (bh *burrowHeap) insert(gc heapNode) {
	*bh = append(*bh, gc)
	i := len(*bh) - 1
	for i > 0 && (*bh)[heapParent(i)].cost > (*bh)[i].cost {
		j := heapParent(i)
		bh.swap(i, j)
		i = j
	}
}

func (bh *burrowHeap) heapify(root int) {
	left := heapLeftChild(root)
	if left > len(*bh)-1 {
		return
	}
	right := heapRightChild(root)
	if right > len(*bh)-1 {
		if (*bh)[left].cost < (*bh)[root].cost {
			bh.swap(left, root)
		}
		return
	} else {
		smaller := left
		if (*bh)[right].cost < (*bh)[left].cost {
			smaller = right
		}
		if (*bh)[smaller].cost >= (*bh)[root].cost {
			return
		}
		bh.swap(smaller, root)
		bh.heapify(smaller)
	}
}

func (bh *burrowHeap) pop() heapNode {
	least := (*bh)[0]
	(*bh)[0] = (*bh)[len(*bh)-1]
	*bh = (*bh)[:len(*bh)-1]
	bh.heapify(0)
	return least
}
