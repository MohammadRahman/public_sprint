package sprint

var RomanValueMap = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func FromRoman(s string) int {
	total := 0
	prev := 0

	for i := len(s) - 1; i >= 0; i-- {
		current := RomanValueMap[s[i]]

		if current < prev {
			total -= current
		} else {
			total += current
		}

		prev = current
	}

	return total
}
