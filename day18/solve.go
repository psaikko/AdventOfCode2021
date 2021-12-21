package day18

import (
	"adventofcode/common"
	"fmt"
	"os"
	"strconv"
)

type snailNumber struct {
	value       *int
	left, right *snailNumber
}

func parse(from int, s string, depth string) (int, snailNumber) {

	if s[from] >= '0' && s[from] <= '9' {
		val, _ := strconv.Atoi(string(s[from]))
		return from + 1, snailNumber{
			value: &val,
			left:  nil,
			right: nil,
		}
	}

	i, left := parse(from+1, s, depth+" ")
	j, right := parse(i+1, s, depth+" ")
	return j + 1, snailNumber{
		value: nil,
		left:  &left,
		right: &right,
	}
}

func (sn *snailNumber) String() string {
	if sn.value != nil {
		return strconv.Itoa(*sn.value)
	} else {
		return fmt.Sprintf("[%s,%s]", sn.left, sn.right)
	}
}

func (sn *snailNumber) tryExplode() bool {

	numQueue := []*snailNumber{sn}
	depthQueue := []int{0}
	var explodedNumber *snailNumber = nil
	inorderPtrs := []*int{}
	explodedIndex := -1

	for len(numQueue) > 0 {
		num := numQueue[len(numQueue)-1]
		depth := depthQueue[len(depthQueue)-1]
		numQueue = numQueue[:len(numQueue)-1]
		depthQueue = depthQueue[:len(depthQueue)-1]

		if depth == 4 && num.value == nil && explodedNumber == nil {
			inorderPtrs = append(inorderPtrs, nil)
			explodedNumber = num
			explodedIndex = len(inorderPtrs) - 1
		} else if num.value != nil {
			inorderPtrs = append(inorderPtrs, num.value)
		} else {
			numQueue = append(numQueue, num.right, num.left)
			depthQueue = append(depthQueue, depth+1, depth+1)
		}
	}

	if explodedNumber != nil {
		if explodedIndex > 0 {
			*inorderPtrs[explodedIndex-1] += *explodedNumber.left.value
		}
		if explodedIndex < len(inorderPtrs)-1 {
			*inorderPtrs[explodedIndex+1] += *explodedNumber.right.value
		}
		zero := 0
		*explodedNumber = snailNumber{&zero, nil, nil}
		return true
	}

	return false
}

func (sn *snailNumber) trySplit() bool {

	numQueue := []*snailNumber{sn}
	for len(numQueue) > 0 {
		num := numQueue[len(numQueue)-1]
		numQueue = numQueue[:len(numQueue)-1]
		if num.value != nil {
			v := *num.value
			if v >= 10 {
				rvalue := (v + 1) / 2
				lvalue := v - rvalue
				*num = snailNumber{
					nil,
					&snailNumber{&lvalue, nil, nil},
					&snailNumber{&rvalue, nil, nil},
				}
				return true
			}
		} else {
			numQueue = append(numQueue, num.right, num.left)
		}
	}
	return false
}

func (sn *snailNumber) magnitude() int {
	if sn.value != nil {
		return *sn.value
	}
	return 3*sn.left.magnitude() + 2*sn.right.magnitude()
}

func add(a, b *snailNumber) *snailNumber {
	sum := snailNumber{
		nil,
		a,
		b,
	}

	reduced := true
	for reduced {
		reduced = sum.tryExplode()
		if !reduced {
			if sum.trySplit() {
				reduced = true
				continue
			}
		}
	}
	return &sum
}

func Run() {
	lines := common.ReadStringLines(os.Stdin)

	var prev *snailNumber = nil

	snailNumbers := []*snailNumber{}

	for _, line := range lines {
		_, next := parse(0, line, "")
		snailNumbers = append(snailNumbers, &next)

		if prev == nil {
			prev = &next
			fmt.Println(prev)
		} else {
			fmt.Println(prev)
			fmt.Println("+", &next)
			prev = add(prev, &next)
		}
	}
	fmt.Println("Final sum", prev.magnitude())

	maxMag := 0
	for i := range snailNumbers {
		for j := range snailNumbers {
			if i != j {
				_, a := parse(0, lines[i], "")
				_, b := parse(0, lines[j], "")
				mag := add(&a, &b).magnitude()
				maxMag = common.Max(mag, maxMag)
			}
		}
	}
	fmt.Println("Best sum of two", maxMag)
}
