package medium

import (
	"testing"
	"sort"
)

/**
 * https://leetcode.com/problems/house-robber/
 * You are a professional robber planning to rob houses along a street. 
 * Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected and it will automatically contact the police if two adjacent houses were broken into on the same night.
 * Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.
 * Constraints:
 * 1 <= nums.length <= 100
 * 0 <= nums[i] <= 400
 * 
 * soln: https://leetcode.com/problems/house-robber/solutions/156523/From-good-to-great.-How-to-approach-most-of-DP-problems
 */

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	//create slice of nodes
	var nodes []node = make([]node, n)
	for index, val := range nums {
		n := node {
			num:   val,
			index: index,
		}
		nodes[index] = n
	}
	//sort descending
	sort.Slice(nodes, func(i, j int) bool { return nodes[i].num > nodes[j].num })

	var bits []bool = make([]bool, n)
	sum, idx := 0, 0
	for _, node := range nodes {
		idx = node.index
		if bits[idx] {
			continue
		}
		sum += node.num
		bits[idx] = true
		if idx != 0 {
			bits[node.index - 1] = true	
		}
		if idx != n-1 {
			bits[node.index + 1] = true	
		}
	}
	return sum
}

type node struct {
	num   int
	index int
}

func TestRob(t *testing.T) {
	tcases := []struct {
		in []int
		want int
	}{
		{
			in:   []int{1,2,3,1},
			want: 4,
		},
		{
			in:   []int{2,7,9,3,1},
			want: 12,
		},
		{
			in:   []int{2,1,1,2},
			want: 4,
		},
		{
			in:   []int{1,1},
			want: 1,
		},
		{
			in:   []int{1,2,1,3,100},
			want: 102,
		},
	} 
	for _, tc := range tcases {
		if got := rob(tc.in); got != tc.want {
			t.Logf("testcase failed. input: %v, expected: %d, got: %d\n", tc.in, tc.want, got)
			t.Fail()
		}
	}
}
