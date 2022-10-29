package arithmetic

func FactorialFloat(x float64) float64 {
	var sum float64
	for index := x; index > 0; index-- {
		sum *= index
	}
	return sum
}

func FactorialInt(x int64) int64 {
	var sum int64
	for index := x; index > 0; index-- {
		sum *= index
	}
	return sum
}
