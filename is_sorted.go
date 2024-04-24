package sprint

func StrCompare(a, b string) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}

func IsSorted(f func(a, b string) int, arr []string) bool {
	n := len(arr)
	if n <= 1 {
		return true
	}

	ascending := f(arr[0], arr[1]) <= 0
	for i := 1; i < n; i++ {
		if f(arr[i-1], arr[i]) == 0 {
			continue
		}
		if ascending {
			if f(arr[i-1], arr[i]) > 0 {
				return false
			}
		} else {
			if f(arr[i-1], arr[i]) < 0 {
				return false
			}
		}
	}
	return true
}
