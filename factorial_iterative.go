package sprint

func FactorialIterative(n int) int {
	if n < 0 {
		return 0
	}

	// Initialize the result to 1
	result := 1

	// Compute factorial iteratively
	for i := 1; i <= n; i++ {
		if result > 0 && i > (1<<31-1)/result {
			return 0
		}
		result *= i
	}
	return result
}
