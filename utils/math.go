package utils

func MaxI(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func MinI(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func ClampI(value, a, b int) int {
	if value < a {
		return a
	} else if value > b {
		return b
	}

	return value
}
