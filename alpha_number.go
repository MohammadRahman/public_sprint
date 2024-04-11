package sprint

func AlphaNumber(n int) string {
	var result string
	negative := n < 0
	if negative {
		n = -n
	} else if n == 0 {
		return "a"
	}
	digitToLetter := "abcdefghij"

	for n > 0 {
		digit := n % 10
		letter := digitToLetter[digit]
		result = string(letter) + result
		n = n / 10
	}
	if negative {
		result = "-" + result
	}

	return result
}
