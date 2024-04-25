package sprint

import "fmt"

func StrCompress(input string) string {
	if len(input) <= 1 {
		return input
	}

	var compressed string
	count := 1

	for i := 1; i < len(input); i++ {
		if input[i] == input[i-1] {
			count++
		} else {
			if count > 1 {
				compressed += fmt.Sprintf("%d", count)
			}
			compressed += string(input[i-1])
			count = 1
		}
	}

	if count > 1 {
		compressed += fmt.Sprintf("%d", count)
	}
	compressed += string(input[len(input)-1])

	return compressed
}
