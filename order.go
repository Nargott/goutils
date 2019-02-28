package goutils

// Min returns a minimal value of a and b
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
