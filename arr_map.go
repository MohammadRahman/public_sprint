package sprint

func ArrMap(f func(int) bool, a []int) []bool {
	var result []bool
	for _, char := range a {
		if f(char) {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}
	return result
}
func IsNegative(x int) bool {
	if x < 0 {
		return true
	}
	return false
}

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}

	if n == 2 {
		return true
	}

	if n%2 == 0 {
		return false
	}

	maxDivisor := int(approxSqrt(float64(n)))
	for i := 3; i <= maxDivisor; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func approxSqrt(x float64) float64 {
	if x == 0 {
		return 0
	}
	z := x / 2.0
	for i := 0; i < 20; i++ {
		z = (z + x/z) / 2.0
	}
	return z
}
