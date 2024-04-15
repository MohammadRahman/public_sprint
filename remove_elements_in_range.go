package sprint

func RemoveElementsInRange(arr []float64, from, to int) []float64 {
	// Normalize negative indices
	if from < 0 {
		from = len(arr) + from
	}
	if to < 0 {
		to = len(arr) + to
	}

	if from > to {
		from, to = to, from
	}
	if from < 0 {
		from = 0
	}
	if to >= len(arr) {
		to = len(arr) - 1
	}
	var result []float64
	if from > 0 {
		result = append(result, arr[:from]...)
	}
	if to > 0 {
		result = append(result, arr[to:]...)
	}

	return result
}
