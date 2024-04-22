package sprint

import "math"

func FactorialIterative(n int) int {
	// Check for invalid input
	if n < 0 {
		return 0
	}
	result := 1
	for i := 1; i <= n; i++ {
		if result > math.MaxInt32/i {
			return 0
		}
		result *= i
	}
	return result
}
