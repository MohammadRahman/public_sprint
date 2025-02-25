package sprint

func NextPrime(n int) int {
	if n <= 1 {
		return 2
	}

	for num := n; ; num++ {
		if isPrime(num) {
			return num
		}
	}
}

// Function to check if a number is prime
func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	if num <= 3 {
		return true
	}
	if num%2 == 0 || num%3 == 0 {
		return false
	}

	i := 5
	for i*i <= num {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
		i += 6
	}

	return true
}
