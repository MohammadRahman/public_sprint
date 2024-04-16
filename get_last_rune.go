package sprint

func GetLastRune(s string) rune {
	toRunes := []rune(s)
	lastRune := toRunes[len(toRunes)-1]
	return lastRune
}
