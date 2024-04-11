package sprint

func Countdown(n int) string {
	var result string
	for i := n; i >= 0; i -= 2 {
		digit := '0' + rune(i)
		result += string(digit) + ", "
	}
	return result + "0!"
}
