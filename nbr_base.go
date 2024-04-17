package sprint

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	for _, char := range base {
		if char == '+' || char == '-' {
			return false
		}
	}
	return true
}
func indexOf(r rune, base string) int {
	for i, b := range base {
		if b == r {
			return i
		}
	}
	return -1
}
func NbrBase(n int, base string) string {
	if !isValidBase(base) {
		return "NV" // Not a valid base
	}

	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}

	baseLen := len(base)
	if n == 0 {
		return string(base[0])
	}

	result := ""
	for n > 0 {
		remainder := n % baseLen
		result = string(base[remainder]) + result
		n = n / baseLen
	}

	return sign + result
}
