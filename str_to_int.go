package sprint

func StrToInt(s string) int {
	var num int
	sign := 1
	start := 0

	for start < len(s) && s[start] == ' ' {
		start++
	}
	if start < len(s) && (s[start] == '+' || s[start] == '-') {
		if s[start] == '-' {
			sign = -1
		}
		start++
	}

	for i := start; i < len(s); i++ {
		digit := int(s[i] - '0')
		if digit < 0 || digit > 9 {
			return 0
		}
		num = num*10 + digit
	}
	return sign * num
}
