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
func ConvertDecToAny(n int, base string) string {
	// Validate the base
	if !isValidBase(base) {
		return "0" // Not a valid base, return "0"
	}

	// Handle special case for zero
	if n == 0 {
		return string(base[0]) // Return first character of base
	}

	result := ""
	baseLen := len(base)

	// Convert decimal (base-10) integer n to specified base
	for n > 0 {
		remainder := n % baseLen
		result = string(base[remainder]) + result
		n = n / baseLen
	}

	return result
}
func ConvertAnyToAny(nbr, baseFrom, baseTo string) string {
	// Validate the bases
	if !isValidBase(baseFrom) || !isValidBase(baseTo) {
		return "0"
	}

	// Convert nbr from baseFrom to decimal (base-10) integer
	decimalValue := ConvertAnyToDec(nbr, baseFrom)

	// Convert decimalValue from base-10 to baseTo
	convertedValue := ConvertDecToAny(decimalValue, baseTo)

	return convertedValue
}
