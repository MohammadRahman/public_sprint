package sprint

func PascalsTriangle(n int) [][]int {
	if n <= 0 {
		return [][]int{}
	}
	triangle := [][]int{{1}}

	for i := 1; i < n; i++ {
		prevRow := triangle[i-1]
		currentRow := make([]int, i+1)
		currentRow[0] = 1
		currentRow[i] = 1
		for j := 1; j < i; j++ {
			currentRow[j] = prevRow[j-1] + prevRow[j]
		}
		triangle = append(triangle, currentRow)
	}

	return triangle
}
