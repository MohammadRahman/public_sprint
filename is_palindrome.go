package sprint

func removeSpacesAndToLower(s string) string {
	var result []rune
	for _, char := range s {
		if char != ' ' {
			if 'A' <= char && char <= 'Z' {
				char += 32
			}
			result = append(result, char)
		}
	}
	return string(result)
}
func IsPalindrome(s string) bool {
	s = removeSpacesAndToLower(s)
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
