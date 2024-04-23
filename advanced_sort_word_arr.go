package sprint

func AdvancedSortWordArr(a []string, f func(a, b string) int) []string {
	n := len(a)
	if n <= 1 {
		return a
	}

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if f(a[j], a[j+1]) > 0 {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}

func StrCompare(a, b string) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}
