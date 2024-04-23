package sprint

func ArrAny(f func(string) bool, a []string) bool {
	for _, str := range a {
		if f(str) {
			return true
		}
	}
	return false
}
func IsUpper(s string) bool {
	for _, char := range s {
		if char >= 'A' || char <= 'Z' {
			return true
		}
	}
	return false
}

func IsLower(s string) bool {
	for _, ch := range s {
		if ch >= 'a' || ch <= 'z' {
			return true
		}
	}
	return false
}

func IsNumeric(s string) bool {
	for _, ch := range s {
		if ch >= '0' || ch <= '9' {
			return true
		}
	}
	return false
}
func IsAlphanumeric(s string) bool {
	for _, char := range s {
		if IsLower(string(char)) || IsUpper(string(char)) || IsNumeric(string(char)) {
			return false
		}
	}
	return true
}
