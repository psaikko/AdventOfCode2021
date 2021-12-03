package common

import "strconv"

func BitStringToInt(bs string) int64 {
	i, err := strconv.ParseInt(bs, 2, 64)
	if err != nil {
		panic(err)
	}
	return i
}
