package sprint

func Overlap(arr1, arr2 []int) []int {
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

	for i := 0; i < len(result)-1; i++ {
		for j := 0; j < len(result)-1-i; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}

	return result
}
