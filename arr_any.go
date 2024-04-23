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
		if 'A' <= char && char <= 'Z' {
			continue
		}
		return false
	}
	return true
}
func IsAlphanumeric(s string) bool {
	for _, char := range s {
		if ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z') || ('0' <= char && char <= '9') {
			continue
		}
		return false
	}
	return true
}
