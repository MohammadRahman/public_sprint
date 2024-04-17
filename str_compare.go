package sprint

func StrCompare(a, b string) int {
	minLength := min(len(a), len(b))

	for i := 0; i < minLength; i++ {
		if a[i] < b[i] {
			return -1
		} else if a[i] > b[i] {
			return 1
		}
	}
	if len(a) < len(b) {
		return -1
	} else if len(a) > len(b) {
		return 1
	}
	return 0
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
