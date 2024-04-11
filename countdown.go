package main

import "fmt"

func Countdown(n int) string {
	if n < 0 || n > 9 {
		return "" // Return empty string for invalid input
	}

	var result string

	// Iterate from n down to 0 (inclusive), skipping every second number
	for i := n; i >= 0; i -= 2 {
		// Convert the current number to its string representation
		currentStr := i
		// Append the current number to the result string
		if len(result) > 0 {
			result += ", "
		}
		result += currentStr
	}

	// Append the ending string "0!" to the result
	result += ", 0!"

	return result
}

func main() {
	c := Countdown(6)
	fmt.Println(c)
}
