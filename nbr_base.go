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
func NbrBase(n int, base string) string {
	if !isValidBase(base) {
		return "NV"
	}

	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}

	if n == 0 {
		return base[0:1]
	}

	result := ""
	baseLen := len(base)

	for n > 0 {
		remainder := n % baseLen
		result = string(base[remainder]) + result
		n = n / baseLen
	}

	return sign + result
}
