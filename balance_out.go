package sprint

func BalanceOut(arr []bool) []bool {

	trueCount := 0
	falseCount := 0

	for _, value := range arr {
		if value {
			trueCount++
		} else {
			falseCount++
		}
	}
	// we found 3false count and 1 true count
	var needed bool

	diff := abs(trueCount - falseCount)

	if trueCount > falseCount {
		needed = false
	} else {
		needed = true
	}

	for i := 0; i < diff; i++ {
		arr = append(arr, needed)
	}
	return arr
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
