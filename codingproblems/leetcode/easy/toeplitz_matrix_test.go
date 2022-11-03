package easy

import (
	"testing"
)

/*
https://leetcode.com/problems/toeplitz-matrix/
Given an m x n matrix, return true if the matrix is Toeplitz. Otherwise, return false.
A matrix is Toeplitz if every diagonal from top-left to bottom-right has the same elements.
*/
func isToeplitzMatrix(matrix [][]int) bool {
	m := len(matrix)
	n := len(matrix[0])
	if m == 1 && n >= 1 || n == 1 && m >= 1 {
		return true
	}
	var p1, p2, jEnd, p1End, p2End int = 0, 1, n - 2, m - 2, m - 1
	for {
		for j := 0; j <= jEnd; j++ {
			if matrix[p1][j] != matrix[p2][j+1] {
				return false
			}
		}
		p1++
		p2++
		if p1 > p1End || p2 > p2End {
			return true
		}
	}
}

func isToeplitzMatrixSol1(matrix [][]int) bool {
	rows := len(matrix) - 1
	cols := len(matrix[0]) - 1

	row := 1
	for row <= rows {
		col := 1
		for col <= cols {
			if matrix[row][col] != matrix[row-1][col-1] {
				return false
			}
			col = col + 1
		}
		row = row + 1
	}
	return true
}

func TestIsToeplitzMatrix(t *testing.T) {
	tcs := []struct {
		in   [][]int
		want bool
	}{
		{
			[][]int{{1, 2, 3, 5, 6}, {4, 1, 2, 3, 5}, {6, 4, 1, 2, 3}, {9, 6, 4, 1, 2}},
			true,
		},
		{
			[][]int{{1, 2}, {2, 2}},
			false,
		},
		{
			[][]int{{4, 3}, {2, 2}},
			false,
		},
		{
			[][]int{{4, 3}},
			true,
		},
		{
			[][]int{{4}, {3}},
			true,
		},
	}

	for _, tc := range tcs {
		got := isToeplitzMatrix(tc.in)
		if got != tc.want {
			t.Logf("Test cased failed. inp: %v got: %t want: %t \n", tc.in, got, tc.want)
		}
	}
}

var gres bool

func BenchmarkIsToeplitzMatrix(b *testing.B) {
	matrix := [][]int{{1, 2, 3, 5, 6}, {4, 1, 2, 3, 5}, {6, 4, 1, 2, 3}, {9, 6, 4, 1, 2}}
	var res bool
	for i := 0; i < b.N; i++ {
		res = isToeplitzMatrix(matrix)
	}
	gres = res
}

func BenchmarkIsToeplitzMatrixSol1(b *testing.B) {
	matrix := [][]int{{1, 2, 3, 5, 6}, {4, 1, 2, 3, 5}, {6, 4, 1, 2, 3}, {9, 6, 4, 1, 2}}
	var res bool
	for i := 0; i < b.N; i++ {
		res = isToeplitzMatrixSol1(matrix)
	}
	gres = res
}
