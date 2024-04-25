package sprint

func Overlap(arr1, arr2 []int) []int {
	if len(arr1) == 0 || len(arr2) == 0 {
		return []int{}
	}

	count := make(map[int]int)
	for _, num := range arr1 {
		count[num]++
	}

	var result []int
	for _, num := range arr2 {
		if count[num] > 0 {
			result = append(result, num)
			count[num]--
		}
	}
	return result
}
