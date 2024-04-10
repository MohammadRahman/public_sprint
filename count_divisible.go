package sprint

func CountDivisible(from, to, step, divisior int) int {
	if step <= 0 || divisior == 0 {
		return 0
	}
	count := 0
	for i := from; i < to; i += step {
		if i%divisior == 0 {
			count++
		}
	}
	return count
}
