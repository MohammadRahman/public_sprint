package sprint

// ch == 65 - 87
func isUpper(s string) bool {
	for _, ch := range s {
		if ch >= 'A' || ch <= 'Z' {
			return true
		}
	}
	return false
}
func isAlphaNumeric(s string) bool {
	for _, ch := range s {
		if (isUpper(string(ch)) || ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9') {
			return true
		}
	}
	return false
}
func ArrAny(f func(string) bool, a []string) bool {
	for _, ch := range a {
		if f(ch) {
			return true
		}
	}
	return false
}
