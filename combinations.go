package main

import "fmt"

func Combinations() string {
	var result string

	for i := 0; i < 10; i++ {
		for j := i + 1; j < 10; j++ {
			for k := j + 1; k < 10; k++ { // Fixed k increment here
				// Append the triplet to the result string
				if i < j && j < k {

					result += fmt.Sprintf("%d%d%d, ", i, j, k)
				}
			}
		}
	}
	if len(result) > 0 {
		result = result[:len(result)-2]
	}

	return result
}

func main() {
	c := Combinations()
	fmt.Println(c)
}
