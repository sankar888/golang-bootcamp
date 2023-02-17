package dynamicprogram

import (
	"testing"
)

/**
 * https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
 * You are given an array prices where prices[i] is the price of a given stock on the ith day.
 * You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
 * Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
 *
 * Constraints:
 * 1 <= prices.length <= 10^5
 * 0 <= prices[i] <= 10^4
 */
func maxProfit(prices []int) int {
	var len int = len(prices)
	if len == 1 {
		return 0
	}
	var greatest, profit int = prices[len-1], 0
	for i := len - 2; i >= 0; i-- {
		if diff := greatest - prices[i]; diff > 0 {
			if diff > profit {
				profit = diff
			}
		} else {
			greatest = prices[i]
		}
	}
	return profit
}

func TestMaxProfit(t *testing.T) {
	tcases := []struct {
		in   []int
		want int
	}{
		{
			in:   []int{7, 1, 5, 3, 6, 4},
			want: 5,
		},
		{
			in:   []int{7, 6, 4, 3, 1},
			want: 0,
		},
		{
			in:   []int{1},
			want: 0,
		},
		{
			in:   []int{2, 4, 3, 4, 2},
			want: 2,
		},
		{
			in:   []int{6, 6, 6, 6, 6},
			want: 0,
		},
	}
	for _, tc := range tcases {
		if got := maxProfit(tc.in); got != tc.want {
			t.Logf("testcase failed. input: %v, expected: %d, got: %d\n", tc.in, tc.want, got)
			t.Fail()
		}
	}
}
