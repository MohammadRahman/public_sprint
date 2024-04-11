package sprint

func SimpleStrToInt(s string) int {
	var num int

	for _, char := range s {
		digit := char - '0'

		if digit < 0 || digit > 9 {
			return 0
		}
		num = num*10 + int(digit)
	}
	return num
}
