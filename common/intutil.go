package common

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func Max(vals ...int) int {
	max := vals[0]
	for _, v := range vals[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func Min(vals ...int) int {
	min := vals[0]
	for _, v := range vals[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

func Sum(vals ...int) int {
	sum := 0
	for _, v := range vals {
		sum += v
	}
	return sum
}

func Prod(vals ...int) int {
	prod := 1
	for _, v := range vals {
		prod *= v
	}
	return prod
}
