package sprint

func ToRoman(num int) string {
	if num < 1 || num > 3999 {
		return "Invalid input"
	}
	romanSymbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	romanValues := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	result := ""

	for i := 0; num > 0; i++ {
		for num >= romanValues[i] {
			result += romanSymbols[i]
			num -= romanValues[i]
		}
	}

	return result
}
