package sprint

func ReverseAlphabet(step int) string {
	const startLetter = 'z'
	if step <= 0 {
		step = 1
	}
	var result string
	for currentLetter := startLetter; currentLetter >= 'a'; currentLetter-- {
		result += string(currentLetter)
		nextLetter := currentLetter - rune(step-1)
		if nextLetter < 'a' {
			break
		}
		currentLetter = nextLetter
	}
	return result
}
