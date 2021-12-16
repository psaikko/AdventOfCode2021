package day16

import (
	"adventofcode/common"
	"fmt"
	"os"
	"strconv"
)

type packet struct {
	bits string
	read int
}

func (p *packet) readInt(next int) int {
	i, _ := strconv.ParseInt(p.bits[p.read:p.read+next], 2, 64)
	p.read += next
	return int(i)
}

func (p *packet) parse() (int, int) {
	version := p.readInt(3)
	packetType := p.readInt(3)

	if packetType == 4 {
		data := 0
		continueBit := 1
		for continueBit == 1 {
			continueBit = p.readInt(1)
			chunk := p.readInt(4)
			data = data<<4 + chunk
		}
		return version, data
	} else {
		lengthTypeID := p.readInt(1)
		subpacketsData := []int{}
		if lengthTypeID == 0 {
			subpacketsLength := p.readInt(15)
			subpacketsStart := p.read
			for p.read < subpacketsStart+subpacketsLength {
				subpacketSum, subpacketData := p.parse()
				subpacketsData = append(subpacketsData, subpacketData)
				version += subpacketSum
			}
		} else {
			subpacketsCount := p.readInt(11)
			for i := 0; i < int(subpacketsCount); i++ {
				subpacketSum, subpacketData := p.parse()
				subpacketsData = append(subpacketsData, subpacketData)
				version += subpacketSum
			}
		}

		data := 0
		switch packetType {
		case 0:
			data = common.Sum(subpacketsData...)
		case 1:
			data = common.Prod(subpacketsData...)
		case 2:
			data = common.Min(subpacketsData...)
		case 3:
			data = common.Max(subpacketsData...)
		case 5:
			if subpacketsData[0] > subpacketsData[1] {
				data = 1
			}
		case 6:
			if subpacketsData[0] < subpacketsData[1] {
				data = 1
			}
		case 7:
			if subpacketsData[0] == subpacketsData[1] {
				data = 1
			}
		}
		return version, data
	}
}

func decode(s string) (int, int) {
	bitString := ""
	for _, r := range s {
		val, _ := strconv.ParseInt(string(r), 16, 64)
		bitString += fmt.Sprintf("%04b", val)
	}
	p := packet{bitString, 0}
	return p.parse()
}

func Run() {
	lines := common.ReadStringLines(os.Stdin)
	for _, line := range lines {
		fmt.Println(decode(line))
	}
}
