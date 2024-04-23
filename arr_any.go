package sprint

func ArrAny(f func(string) bool, a []string) bool {
	for _, str := range a {
		if f(str) {
			return true
		}
	}
	return false
}

func IsLower(s string) bool {
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			return true
		}
	}
	return false
}
func IsNumeric(s string) bool {
	for _, char := range s {
		if char >= '0' && char <= '9' {
			return true
		}
	}
	return false
}
