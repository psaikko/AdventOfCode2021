package common

import (
	"bufio"
	"io"
	"strconv"
)

// ReadIntLines reads single-integer lines from an io.Reader
func ReadIntLines(r io.Reader) []int {
	lines := make([]int, 0, 100)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		lines = append(lines, val)
	}

	return lines
}
