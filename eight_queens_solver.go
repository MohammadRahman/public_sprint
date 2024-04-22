package sprint

import (
	"strconv"
	"strings"
)

func EightQueensSolver() string {
	var solutions []string
	board := make([]int, 8)
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
			var solution strings.Builder
			for _, row := range board {
				solution.WriteString(strconv.Itoa(row + 1))
			}
			solutions = append(solutions, solution.String())
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
	var result strings.Builder
	for i, sol := range solutions {
		result.WriteString(sol)
		if i < len(solutions)-1 {
			result.WriteString("\n")
		}
	}

	return result.String()
}

func absolute(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
