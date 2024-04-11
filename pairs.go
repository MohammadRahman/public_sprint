package sprint

import "fmt"

func Pairs() string {
	var result string

	for i := 0; i <= 99; i++ {
		for j := i + 1; j <= 99; j++ {
			pair := fmt.Sprintf("%02d %02d", i, j)

			if len(result) > 0 {
				result += ", "
			}
			result += pair
		}
	}
	return result
}
