package sprint

func SortIntegerTable(table []int) []int {
	n := len(table)
	if n <= 1 {
		return table // If the slice has 0 or 1 element, it's already sorted
	}
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if table[j] > table[j+1] {
				table[j], table[j+1] = table[j+1], table[j]
				swapped = true
			}
		}
		if !swapped {
			break // If no swaps are made, the list is already sorted
		}
	}
	return table
}
