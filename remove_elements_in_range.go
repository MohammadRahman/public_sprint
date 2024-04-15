package main

import "fmt"

func RemoveElementsInRange(arr []float64, from, to int) []float64 {
	// Normalize negative indices
	if from < 0 {
		from = len(arr) + from
	}
	if to < 0 {
		to = len(arr) + to
	}

	// Ensure from <= to
	if from > to {
		from, to = to, from
	}

	// Validate index boundaries
	if from < 0 || to >= len(arr) || from >= to {
		return arr
	}

	// Create a new slice to store the filtered elements
	var result []float64

	// Append elements before the 'from' index
	result = append(result, arr[:from]...)

	// Append the element at the 'to' index
	result = append(result, arr[to])

	return result

}

func main() {
	fmt.Println(RemoveElementsInRange([]float64{10., .8, -.4, 20., 7.7, 3.}, 4, 1))
}
