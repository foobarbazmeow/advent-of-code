package util

func SumInts(xs []int) int {
	result := 0
	for _, v := range xs {
		result += v
	}
	return result
}
