package sprint

func IsSorted(f func(a, b string) int, arr []string) bool {
	n := len(arr)
	if n <= 1 {
		return true
	}
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if f(arr[i-1], arr[i]) > 0 {
				return false
			}
		}
	}
	return true
}

func StrCompare(a, b string) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	}
	return -1
}
