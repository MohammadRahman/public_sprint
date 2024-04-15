package sprint

import "fmt"

func CombN(n int) []string {
	var result []string
	generateCombinations("", n, &result)
	return result
}

// Recursive helper function to generate combinations
func generateCombinations(current string, n int, result *[]string) {
	if n == 0 {
		// Append the current combination to the result
		*result = append(*result, current)
		return
	}

	start := 0
	if len(current) > 0 {
		// Start with the next digit after the last digit in the current combination
		start = int(current[len(current)-1]-'0') + 1
	}

	for i := start; i <= 9; i++ {
		// Recursively generate combinations with the current digit appended
		generateCombinations(current+fmt.Sprintf("%d", i), n-1, result)
	}
}
