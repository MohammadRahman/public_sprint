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
func ConvertAnyToDec(s string, base string) int {
	// Validate the base
	if !isValidBase(base) {
		return 0 // Not a valid base, return 0
	}

	// Map to store character positions for base values
	baseMap := make(map[rune]int)
	for i, char := range base {
		baseMap[char] = i
	}

	decimalValue := 0
	baseLen := len(base)

	// Convert s to decimal (base-10) integer
	for _, char := range s {
		if val, ok := baseMap[char]; ok {
			decimalValue = decimalValue*baseLen + val
		} else {
			return 0 // Invalid character found in s, return 0
		}
	}

	return decimalValue
}
