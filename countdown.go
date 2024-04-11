package sprint

func Countdown(n int) string {
	var result string
	for i := n; i >= 0; i -= 2 {
		digit := '0' + rune(i)
		result += string(digit)
		if i > 0 {
			result += ", "
		}
	}
	if result[len(result)-1] == '0' {
		result += "!" // If last character is '0', append "!"
	} else {
		result += "0!" // Otherwise, append "0!"
	}
	return result
}
