package dynamicprogram

import (
	"testing"
	"fmt"
	"sort"
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
	length := len(prices)
	if length == 1 {
		return 0
	}

	// check if all day prices are equal [6,6,6,6,6]
	// check if all day prices goes descending [6,5,4,3,2]
	descending := true
	for i := 1; i < length; i++ {
		if prices[i-1] < prices[i] {
			descending = false
			break
		}
	}
	
	allSame := true
	for i := 1; i < length; i++ {
		if prices[i-1] != prices[i] {
			allSame = false
			break
		}
	}

	if descending || allSame {
		return 0
	}

	//create []dayprice
	var dayprices []dayprice = make([]dayprice, length)
	for i := 0; i < length; i++ {
		dayprices[i] = dayprice {
			price: prices[i],
			day:   i,
		}
	}

	//sort dayprices
	fmt.Println("before dayprices", dayprices)
	sort.Slice(dayprices, func (i, j int) bool { return dayprices[i].price > dayprices[j].price})
	fmt.Println("sorted dayprices", dayprices)
	
	//try to find max profit
	profit := 0
	outer:
	for i := 1; i < length; i++ {
		for j := length-1; j >= 1; j-- {
			if dayprices[i-1].day > dayprices[j].day {
				profit = dayprices[i-1].price - dayprices[j].price
				break outer
			}
		}
	}
	return profit

}

type dayprice struct {
	price int
	day   int
}

func TestMaxProfit(t *testing.T) {
	tcases := []struct {
		in   []int
		want int
	}{
		{
			in:   []int{7,1,5,3,6,4},
			want: 5,
		},
		{
			in:   []int{7,6,4,3,1},
			want: 0,
		},				
		{
			in:   []int{1},
			want: 0,
		},		
		{
			in:   []int{2,4,3,4,2},
			want: 2,
		},
		{
			in:   []int{6,6,6,6,6},
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