package sprint

func SortWordArr(a []string) []string {
	n := len(a)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if compareStrings(a[j], a[j+1]) > 0 {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

	return a
}

func compareStrings(s1, s2 string) int {
	minLen := len(s1)
	if len(s2) < minLen {
		minLen = len(s2)
	}

	for i := 0; i < minLen; i++ {
		if s1[i] != s2[i] {
			return int(s1[i]) - int(s2[i])
		}
	}
	return len(s1) - len(s2)
}
