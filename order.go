package goutils

// Min returns a smallest value of a and b
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Max returns a biggest value of a and b
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
