package main

import "fmt"

func Countdown(n int) string {
	if n < 0 || n > 9 {
		return "" // Return empty string for invalid input
	}

	var result string

	// Iterate from n down to 0 (inclusive), skipping every second number
	for i := n; i >= 0; i -= 2 {
		// Append the current number to the result string
		if len(result) > 0 {
			result += ", "
		}
		result += fmt.Sprintf("%d", i)

		// Check if i is 0 to append the "0!" ending
		if i == 0 {
			result += "!"
		} else if i == 1 {
			// If i is 1 and we haven't reached 0 yet, append "0!" to end the sequence
			result += ", 0!"
			break
		}
	}

	return result
}

func main() {
	c := Countdown(7)
	fmt.Println(c)
}
