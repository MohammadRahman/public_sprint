package sprint

func RemoveElementsInRange(arr []float64, from, to int) []float64 {

	switch {
	case from < 0:
		from = 0
	case from >= len(arr):
		from = len(arr)
	}
	switch {
	case to < 0:
		to = 0
	case to >= len(arr):
		to = len(arr)
	}
	if from > to {
		from, to = to, from
	}
	if from < 0 && to < 0 {
		return arr
	}
	return append(arr[:from], arr[to:]...)
}
