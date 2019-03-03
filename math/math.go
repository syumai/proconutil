package math

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Max(n ...int) int {
	lenN := len(n)
	if lenN == 1 {
		return n[0]
	}
	result := n[0]
	for i := 1; i < lenN; i++ {
		result = max(result, n[i])
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Min(n ...int) int {
	lenN := len(n)
	if lenN == 1 {
		return n[0]
	}
	result := n[0]
	for i := 1; i < lenN; i++ {
		result = min(result, n[i])
	}
	return result
}

// GCD is a GCD function implemented in Euclidean Algorithm
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
