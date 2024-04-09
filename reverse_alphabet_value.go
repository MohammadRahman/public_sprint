package sprint

func ReverseAlphabetValue(ch rune) rune {
	charCode := int(ch)
	reversedCharCode := 'a' + ('z' - charCode)
	return rune(reversedCharCode)
}
