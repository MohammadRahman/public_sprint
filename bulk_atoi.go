package sprint

func BulkAtoi(arr []string) []int {
	var result []int

	for _, str := range arr {
		num := stringToInt(str)
		result = append(result, num)
	}

	return result
}
func stringToInt(s string) int {
	if s == "" {
		return 0
	}

	var sign int = 1
	if s[0] == '-' {
		sign = -1
		s = s[1:] // Remove the '-' sign from the string
	}

	var result int
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			result = result*10 + int(ch-'0')
		} else {
			// Handle invalid input: non-numeric characters
			return 0
		}
	}

	return result * sign
}
