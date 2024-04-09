package sprint

func ShiftBy(r rune, step int) rune {

	intValue := int(r - 'a')

	shiftValue := (intValue + step) % 26

	if shiftValue < 0 {
		shiftValue += 26
	}
	return rune('a' + shiftValue)
}
