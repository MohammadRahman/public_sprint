package sprint

func BetweenLimits(from, to rune) string {

	start, end := OrderRunes(from, to)

	result := ""

	for r := start + 1; r < end; r++ {
		result += string(r)
	}
	return result
}

func OrderRunes(a, b rune) (rune, rune) {
	switch {
	case a < b:
		return a, b
	case a > b:
		return b, a
	default:
		return a, b
	}
}
