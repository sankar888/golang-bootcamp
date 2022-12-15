package medium

import (
	"testing"
	"fmt"
)
/*
 * https://leetcode.com/problems/minimum-falling-path-sum/
 * Given an n x n array of integers matrix, return the minimum sum of any falling path through matrix.
 * A falling path starts at any element in the first row and chooses the element in the next row that is either directly below or diagonally left/right.
 * Specifically, the next element from position (row, col) will be (row + 1, col - 1), (row + 1, col), or (row + 1, col + 1).
 * 
 * Constraints:
 *  n == matrix.length == matrix[i].length
 *  1 <= n <= 100
 *  -100 <= matrix[i][j] <= 100
 */

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	if n == 1 {
		return matrix[0][0]
	}

	var sumMatrix [][]int = make([][]int, n)
	
	for i := n-1; i >= 0; i-- {
		if i + 1 == n { //last row
			sumMatrix[i] = matrix[i]
			continue
		}
		sumMatrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			val := matrix[i][j]
			min := val + sumMatrix[i+1][j]

			if j - 1 >= 0 {
				tmp := val + sumMatrix[i+1][j-1]
				if tmp < min {
					min = tmp
				}
			}

			if j + 1 < n {
				tmp := val + sumMatrix[i+1][j+1]
				if tmp < min {
					min = tmp
				}
			}
			sumMatrix[i][j] = min
		}
	}

	min := sumMatrix[0][0]
	for j := 1; j < n; j++ {
		if tmp := sumMatrix[0][j]; tmp < min {
			min = tmp
		}
	}
	return min
}

func minFallingPathSumOptimized(matrix [][]int) int {
	n := len(matrix)
	if n == 1 {
		return matrix[0][0]
	}
	
	for i := n-2; i >= 0; i-- {
		for j := 0; j < n; j++ {
			min := matrix[i+1][j]
			if j - 1 >= 0 && matrix[i+1][j-1] < min {
				min = matrix[i+1][j-1]
			}

			if j + 1 < n && matrix[i+1][j+1] < min {
				min = matrix[i+1][j+1]
				
			}
			matrix[i][j] = min + matrix[i][j]
		}
	}
	min := matrix[0][0]
	for j := 1; j < n; j++ {
		if matrix[0][j] < min {
			min = matrix[0][j]
		}
	}
	return min
}

func TestMinFallingPathSum(t *testing.T) {
	matrix := [][]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
	sum := minFallingPathSum(matrix)
	t.Log("minsum", sum)
}

func BenchmarkMinFallingPathSumOptimized(b *testing.B) {
	matrix := [][]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
	sum := 0
	for i := 0; i < b.N; i++ {
		sum = minFallingPathSumOptimized(matrix)
	}
	fmt.Println(sum)
}

func BenchmarkMinFallingPathSum(b *testing.B) {
	matrix := [][]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
	sum := 0
	for i := 0; i < b.N; i++ {
		sum = minFallingPathSum(matrix)
	}
	fmt.Println(sum)
}





