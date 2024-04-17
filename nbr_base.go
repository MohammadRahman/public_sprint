package sprint

func NbrBase(n int, base string) string {
	if !isValidBase(base) {
		return "NV"
	}

	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}

	converted := convertToBase(n, base)

	return sign + converted
}

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	for i := 0; i < len(base); i++ {
		if base[i] == '+' || base[i] == '-' {
			return false
		}
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return false
			}
		}
	}
	return true
}

func convertToBase(n int, base string) string {
	if n == 0 {
		return string(base[0])
	}
	converted := ""
	baseLen := len(base)
	for n > 0 {
		digit := n % baseLen
		converted = string(base[digit]) + converted
		n /= baseLen
	}
	return converted
}
