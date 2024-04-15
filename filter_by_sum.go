package sprint

func FilterBySum(arr [][]int, limit int) [][]int {
	// Create a result slice to store filtered subarrays
	var result [][]int

	// Iterate through each subarray in the 2D array
	for _, subarray := range arr {
		// Calculate the sum of elements in the current subarray
		sum := 0
		for _, num := range subarray {
			sum += num
		}

		// Check if the sum of the current subarray is greater than or equal to the limit
		if sum >= limit {
			// Append the current subarray to the result slice
			result = append(result, subarray)
		}
	}
	if len(result) == 0 {
		return [][]int{} // Return an empty 2D array
	}

	return result
}
