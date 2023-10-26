package method

//泛型的使用

func GetMax_i(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func GetMax_f(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
func GetMax_[T interface{ int | float64 }](a, b T) T {
	if a > b {
		return a
	}
	return b
}
