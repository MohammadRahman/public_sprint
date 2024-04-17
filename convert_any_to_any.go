package sprint

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false // Base must contain at least 2 unique characters
	}
	for _, char := range base {
		if char == '+' || char == '-' {
			return false // Base cannot contain '+' or '-'
		}
	}
	return true
}
func ConvertDecToAny(n int, base string) string {
	if !isValidBase(base) {
		return "0"
	}

	if n == 0 {
		return string(base[0])
	}

	result := ""
	baseLen := len(base)
	for n > 0 {
		remainder := n % baseLen
		result = string(base[remainder]) + result
		n = n / baseLen
	}

	return result
}
func ConvertAnyToAny(nbr, baseFrom, baseTo string) string {
	if !isValidBase(baseFrom) || !isValidBase(baseTo) {
		return "0"
	}

	decimalValue := ConvertAnyToDec(nbr, baseFrom)

	convertedValue := ConvertDecToAny(decimalValue, baseTo)

	return convertedValue
}
