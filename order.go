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

// MaxUint64 returns a biggest value of a and b byt uint64 type
func MaxUint64(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}
