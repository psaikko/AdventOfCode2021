package template

import (
	"adventofcode/common"
	"fmt"
	"os"
)

func Run() {
	lines := common.ReadStringLines(os.Stdin)
	fmt.Println(lines)
}
