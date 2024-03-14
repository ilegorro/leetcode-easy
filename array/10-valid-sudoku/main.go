package main

import (
	"fmt"
)

func isValidSudoku(board [][]byte) bool {
	dot := "."[0]
	mSquares := make([][]map[byte]struct{}, 3)
	for i := 0; i < 3; i++ {
		mSquares[i] = make([]map[byte]struct{}, 3)
		for j := 0; j < 3; j++ {
			mSquares[i][j] = make(map[byte]struct{}, 0)
		}
	}
	for i := 0; i < 9; i++ {
		mRow := make(map[byte]struct{}, 0)
		mCol := make(map[byte]struct{}, 0)
		for j := 0; j < 9; j++ {
			s1 := board[i][j]
			s2 := board[j][i]
			if _, ok := mRow[s1]; ok && s1 != dot {
				return false
			}
			if _, ok := mCol[s2]; ok && s2 != dot {
				return false
			}
			mRow[s1] = struct{}{}
			mCol[s2] = struct{}{}

			iSquare := int(float64(i) / 3)
			jSquare := int(float64(j) / 3)
			if _, ok := mSquares[iSquare][jSquare][s1]; ok && s1 != dot {
				return false
			}
			mSquares[iSquare][jSquare][s1] = struct{}{}
		}
	}

	return true
}

func main() {
	boardString := []string{
		"53..7....",
		"6..195...",
		".98....6.",
		"8...6...3",
		"4..8.3..1",
		"7...2...6",
		".6....28.",
		"...419..5",
		"....8..79",
	}
	board := make([][]byte, 9)
	for k, row := range boardString {
		board[k] = make([]byte, 9)
		for i := 0; i < 9; i++ {
			board[k][i] = row[i]
		}
	}

	res := isValidSudoku(board)

	fmt.Println(res)
}
