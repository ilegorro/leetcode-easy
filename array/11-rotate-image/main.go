package main

import (
	"fmt"
)

func rotate(matrix [][]int) {
	n := len(matrix) - 1
	m := len(matrix) / 2
	for i := 0; i < m; i++ {
		for j := i; j < n-i; j++ {
			t := matrix[i][j]
			matrix[i][j] = matrix[n-j][i]
			matrix[n-j][i] = matrix[n-i][n-j]
			matrix[n-i][n-j] = matrix[j][n-i]
			matrix[j][n-i] = t
		}
	}
}

func main() {
	// matrix := [][]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// }

	matrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}

	rotate(matrix)

	fmt.Println(matrix)
}
