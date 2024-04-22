package sprint

func NthFibonacci(index int) int {
	// Handle negative indices
	if index < 0 {
		return -1
	}
	// Base cases for Fibonacci sequence
	if index == 0 {
		return 0
	} else if index == 1 {
		return 1
	}
	return NthFibonacci(index-1) + NthFibonacci(index-2)
}
