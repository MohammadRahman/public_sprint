package sprint

func EightQueensSolver() string {
	var solutions []string
	board := [8]int{}
	isSafe := func(col, row int) bool {
		for i := 0; i < col; i++ {
			if board[i] == row || absolute(board[i]-row) == col-i {
				return false
			}
		}
		return true
	}

	var solveNQueens func(int)
	solveNQueens = func(col int) {
		if col == 8 {
			var solution string
			for i := 0; i < 8; i++ {
				solution += string('1' + byte(board[i]))
			}
			solutions = append(solutions, solution)
			return
		}

		for row := 0; row < 8; row++ {
			if isSafe(col, row) {
				board[col] = row
				solveNQueens(col + 1)
			}
		}
	}

	solveNQueens(0)

	var result string
	for _, sol := range solutions {
		result += sol + "\n"
	}

	return result
}

func absolute(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
