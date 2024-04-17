package sprint

func ToCapitalCase(s string) string {
	var result []rune
	wordStart := true

	for _, char := range s {
		if isAlphanumeric(char) {
			if wordStart {
				if char >= 'a' && char <= 'z' {
					char -= 'a' - 'A'
				}
				result = append(result, char)
				wordStart = false
			} else {
				if char >= 'A' && char <= 'Z' {
					char += 'a' - 'A'
				}
				result = append(result, char)
			}
		} else {
			result = append(result, char)
			wordStart = true
		}
	}

	return string(result)
}
func isAlphanumeric(r rune) bool {
	// Check if the rune is a letter (A-Z or a-z) or a digit (0-9)
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}
