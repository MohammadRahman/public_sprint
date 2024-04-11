package main

import "fmt"

func ReverseAlphabet(step int) string {
	const startLetter = 'z' // Start from 'z' and move backwards
	// If step is non-positive, use a default step of 1
	if step <= 0 {
		step = 1
	}

	var result string

	// Iterate backwards through the alphabet starting from 'z'
	for currentLetter := startLetter; currentLetter >= 'a'; currentLetter-- {
		// Append the current letter to the result string
		result += string(currentLetter)

		// Calculate the next letter position (skip step-1 letters)
		nextLetter := currentLetter - rune(step-1)

		// Ensure we stay within the bounds of 'a' to 'z'
		if nextLetter < 'a' {
			break
		}

		// Update current letter to the next letter position
		currentLetter = nextLetter
	}

	// Return the reversed alphabet sequence as a string
	return result
}

func main() {
	c := ReverseAlphabet(5)
	fmt.Println(c)
}
