package day3

import (
	"adventofcode/common"
	"fmt"
	"os"
)

func countPositionBits(bitstrings []string, pos int) (int, int) {
	zeros := 0
	ones := 0
	for _, line := range bitstrings {
		if line[pos] == '0' {
			zeros += 1
		} else {
			ones += 1
		}
	}
	return zeros, ones
}

func filterByPositionChar(arr []string, pos int, char byte) []string {
	newarr := make([]string, 0)
	for _, line := range arr {
		if line[pos] == char {
			newarr = append(newarr, line)
		}
	}
	return newarr
}

func Run() {
	lines := common.ReadStringLines(os.Stdin)
	m := len(lines[0])

	commonBits := ""
	uncommonBits := ""
	for i := 0; i < m; i++ {
		zeros, ones := countPositionBits(lines, i)
		if zeros > ones {
			commonBits += "0"
			uncommonBits += "1"
		} else {
			commonBits += "1"
			uncommonBits += "0"
		}
	}
	power := common.BitStringToInt(commonBits) * common.BitStringToInt(uncommonBits)
	fmt.Println("Power consumption", power)

	commons := lines
	for i := 0; len(commons) > 1; i++ {
		zeros, ones := countPositionBits(commons, i)
		if zeros > ones {
			commons = filterByPositionChar(commons, i, '0')
		} else {
			commons = filterByPositionChar(commons, i, '1')
		}
	}
	uncommons := lines
	for i := 0; len(uncommons) > 1; i++ {
		zeros, ones := countPositionBits(uncommons, i)
		if zeros > ones {
			uncommons = filterByPositionChar(uncommons, i, '1')
		} else {
			uncommons = filterByPositionChar(uncommons, i, '0')
		}
	}
	rating := common.BitStringToInt(commons[0]) * common.BitStringToInt(uncommons[0])
	fmt.Println("Life support rating", rating)
}
