package sprint

func AlphabetMastery(n int) string {
	if n > 26 {
		return "choose number between 1-26"
	}
	var result string

	for i := 0; i < n; i++ {
		letter := rune('a' + i)
		result += string(letter)
	}
	return result
}
