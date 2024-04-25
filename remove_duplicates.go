package sprint

func RemoveDuplicates(arr []int) []int {
	seen := make(map[int]bool)
	unique := []int{}

	for _, num := range arr {
		if !seen[num] {
			seen[num] = true
			unique = append(unique, num)
		}
	}

	return unique
}
