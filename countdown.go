package sprint

func Countdown(n int) string {
	if n < 0 || n > 9 {
		return "" // Return empty string for invalid input
	}

	var result string

	for i := n; i >= 0; i -= 2 {
		if len(result) > 0 {
			result += ", "
		}
		numStr := ""
		for i > 0 {
			numStr = string('0'+rune(i%10)) + numStr
			i /= 10
		}

		result += numStr
		if i == 0 {
			result += "!"
		} else if i == 1 {
			result += ", 0!"
			break
		}
	}
	return result
}
