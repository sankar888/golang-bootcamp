package dynamicprogram

import (
	"testing"
)

/**
 * https://leetcode.com/problems/pascals-triangle/
 * Given an integer numRows, return the first numRows of Pascal's triangle. 
 * In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:
 * 
 * constraints:
 * 1 <= numRows <= 30
 */

func generate(numRows int) [][]int {
	var res [][]int = make([][]int, numRows)
	return generateMemoized(numRows, res)
}

func generateMemoized(numRows int, cache [][]int) [][]int {
	if numRows == 1 {
		cache[0] = []int{1}
		return cache
	}
	if numRows == 2 {
		cache[0] = []int{1}
		cache[1] = []int{1, 1}
		return cache
	}

	row := make([]int, numRows)
	row[0], row[numRows-1] = 1, 1
	
	for i, pre, prerow := 0, generateMemoized(numRows-1, cache), numRows-2; i < numRows-2; i++ {
		row[i+1] = pre[prerow][i] + pre[prerow][i+1]
	}
	cache[numRows-1] = row
	return cache
}

func TestGenerate(t *testing.T) {
	tcases := []struct {
		in   int
		want [][]int
	}{
		{
			in:   1,
			want: [][]int{{1}},
		},
		{
			in:   3,
			want: [][]int{{1}, {1, 1}, {1,2,1}},
		},				
		{
			in:   5,
			want: [][]int{{1}, {1,1}, {1,2,1}, {1,3,3,1}, {1,4,6,4,1}},
		},		
	}
	for _, tc := range tcases {
		got := generate(tc.in)	
		t.Logf("testcase failed. input: %v, expected: %d, got: %d\n", tc.in, tc.want, got)
	}
}
