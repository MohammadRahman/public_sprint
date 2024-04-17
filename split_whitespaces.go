package sprint

func SplitWhitespaces(s string) []string {
	var words []string
	var currentWord []rune

	for _, char := range s {
		if isWhitespace(char) {
			if len(currentWord) > 0 {
				words = append(words, string(currentWord))
				currentWord = nil
			}
		} else {
			currentWord = append(currentWord, char)
		}
	}
	if len(currentWord) > 0 {
		words = append(words, string(currentWord))
	}

	return words
}

func isWhitespace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n'
}
