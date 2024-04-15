package sprint

func BalanceOut(arr []bool) []bool {
	falseCount := 0
	trueCount := 0

	for _, value := range arr {
		if value {
			trueCount++
		} else {
			falseCount++
		}
	}

	diffs := abs(trueCount - falseCount)
	var needed bool
	var result []bool

	if trueCount > falseCount {
		needed = false
	} else {
		needed = true
	}
	for i := 0; i < diffs; i++ {
		result = append(arr, needed)
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
