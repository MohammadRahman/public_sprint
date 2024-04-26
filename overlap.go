package sprint

import "sort"

func Overlap(arr1, arr2 []int) []int {
	count := make(map[int]int)
	for _, num := range arr1 {
		count[num]++
	}
	common := make([]int, 0)
	for _, num := range arr2 {
		if count[num] > 0 {
			common = append(common, num)
			count[num]--
		}
	}
	sort.Ints(common)

	return common
}
