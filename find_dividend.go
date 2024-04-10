package sprint

func FindDividend(from, to, divisior int) int {

	for i := from; i < to; i++ {
		if i%divisior == 0 {
			return i
		}
	}
	return -1
}
