package common

import (
	"bufio"
	"io"
	"strconv"
	"strings"
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

// ReadIntsLines reads multiple-integer lines from an io.Reader
func ReadIntsLines(r io.Reader) [][]int {
	stringsLines := ReadStringsLines(r)
	intsLines := make([][]int, 0, 100)

	for _, tokens := range stringsLines {
		ints := make([]int, len(tokens))
		for i, token := range tokens {
			v, err := strconv.Atoi(token)
			if err != nil {
				panic(err)
			}
			ints[i] = v
		}
		intsLines = append(intsLines, ints)
	}

	return intsLines
}

// ReadStringsLines returns lines split by whitespace from an io.Reader
func ReadStringsLines(r io.Reader) [][]string {
	stringLines := ReadStringLines(r)
	stringsLines := make([][]string, 0, 100)

	for _, line := range stringLines {
		stringsLines = append(stringsLines, strings.Split(line, " "))
	}

	return stringsLines
}

// ReadStringLines returns all lines from an io.Reader
func ReadStringLines(r io.Reader) []string {
	lines := make([]string, 0, 100)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}
