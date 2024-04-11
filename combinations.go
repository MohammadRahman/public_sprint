package sprint

import "fmt"

func Combinations() string {
	var result string

	for i := 0; i <= 6; i++ {
		for j := i + 1; j <= 8; j++ {
			for k := j + 1; k <= 9; j++ {
				triplet := fmt.Sprintf("%d%d%d", i, j, k)
				// Append the triplet to the result string
				if len(result) > 0 {
					result += ", "
				}
				result += triplet
			}
		}
	}
	return result
}
