package sprint

func FactorialRecursive(n int) int {
	if n < 0 {
		return 0
	}

	// Base case: factorial of 0 is 1
	if n == 0 {
		return 1
	}
	return recursiveFactorial(n)
}

func recursiveFactorial(n int) int {
	if n == 0 {
		return 1
	}
	previousFactorial := recursiveFactorial(n - 1)
	if previousFactorial > 0 && n > (1<<31-1)/previousFactorial {
		return 0
	}
	return previousFactorial * n
}
