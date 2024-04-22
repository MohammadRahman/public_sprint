package sprint

func ToThePowerRecursive(n int, power int) int {
	if power < 0 {
		return 0
	}
	// Base case: power = 0, n^0 = 1
	if power == 0 {
		return 1
	}
	return n * ToThePowerRecursive(n, power-1)
}
