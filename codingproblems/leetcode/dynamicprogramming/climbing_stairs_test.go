package dynamicprogram

import (
	"testing"
)

/**
 * https://leetcode.com/problems/climbing-stairs/
 * You are climbing a staircase. It takes n steps to reach the top.
 * Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
 * Constraints:
 * 1 <= n <= 45
 */

var cache map[int]int = make(map[int]int)
func climbStairs(n int) int {	
	if res, ok := cache[n]; ok {
		return res
	} 
	switch {
	case n <= 0:
		return 0
	case n == 1:
	    return 1
	case n == 2:
		return 2    	
	}
	res := climbStairs(n-1) + climbStairs(n-2)
	cache[n] = res
	return res
}

func TestClimbStairs(t *testing.T) {
	tcases := []struct {
		in   int
		want int
	}{
		{
			in:   2,
			want: 2,
		},
		{
			in:   3,
			want: 3,
		},				
		{
			in:   4,
			want: 5,
		},		
	}
	for _, tc := range tcases {
		if got := climbStairs(tc.in); got != tc.want {
			t.Logf("testcase failed. input: %v, expected: %d, got: %d\n", tc.in, tc.want, got)
			t.Fail()
		}
	}
}

