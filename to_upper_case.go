package sprint

func ToUpperCase(s string) string {
	var result string
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			upperChar := char - 'a' + 'A'
			result += string(upperChar)
		} else {
			result += string(char)
		}
	}

	return result
}
