package medium

import (
	"testing"
)

/*
https://leetcode.com/problems/most-stones-removed-with-same-row-or-column/
TODO: the following code found the leaset no of stones removed. we need most no if stones removed.
*/
func removeStones(stones [][]int) int {
	length := len(stones)
	if length == 1 {
		return 1
	}
	var rbit, cbit [10000]bool = [10000]bool{}, [10000]bool{}
	var count int = 0
	for _, stone := range stones {
		if rIndex, cIndex := stone[0], stone[1]; !rbit[rIndex] && !cbit[cIndex] {
			rbit[rIndex] = true
			cbit[cIndex] = true
		} else {
			count++
		}
	}
	return count
}

func TestRemoveStones(t *testing.T) {
	tcases := []struct {
		stones [][]int
		want   int
	}{
		{
			stones: [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}},
			want:   5,
		},
		{
			stones: [][]int{{0, 0}, {0, 2}, {1, 1}, {2, 0}, {2, 2}},
			want:   3,
		},
		{
			stones: [][]int{{0, 0}},
			want:   1,
		},
	}
	for _, tcase := range tcases {
		if got := removeStones(tcase.stones); got != tcase.want {
			t.Errorf("Tcases failed, tcase: %v, got: %d", tcase, got)
		}
	}
}
