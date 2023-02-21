package dynamicprogram

import (
	"testing"
)

/**
 * https://leetcode.com/problems/jump-game-ii/
 * You are given a 0-indexed array of integers nums of length n. You are initially positioned at nums[0].
 * Each element nums[i] represents the maximum length of a forward jump from index i. In other words, if you are at nums[i], you can jump to any nums[i + j] where:
 * 0 <= j <= nums[i] and
 * i + j < n
 * Return the minimum number of jumps to reach nums[n - 1]. The test cases are generated such that you can reach nums[n - 1].
 *
 * Constraints
 * 1 <= nums.length <= 104
 * 0 <= nums[i] <= 1000
 * It's guaranteed that you can reach nums[n - 1].
 */
func jump(nums []int) int {
	size, marker := len(nums), -1
	if size <= 1 {
		return 0
	}

	//a queue which holds the index
	//-1 is a special queue member to count the depth or jumps
	queue := make(chan int, size)
	defer close(queue)

	//memoization
	memo := make(map[int]bool)
	//put firstindex and marker
	queue <- 0
	queue <- marker
	steps := 0
	for {
		i := <-queue
		if i == marker {
			steps += 1
			queue <- marker
			continue
		}
		num := nums[i]
		if num+i >= size-1 {
			steps += 1
			break
		}
		for j := 1; j <= num; j++ {
			if _, exists := memo[i+j]; !exists {
				queue <- (i + j)
				memo[i+j] = true
			}

		}
	}
	return steps
}

func TestJump(t *testing.T) {
	tcases := []struct {
		in   []int
		want int
	}{
		{
			in:   []int{1},
			want: 0,
		},
		{
			in:   []int{1, 2},
			want: 1,
		},
		{
			in:   []int{2, 2, 1},
			want: 1,
		},
		{
			in:   []int{2, 1, 3, 5},
			want: 2,
		},
		{
			in:   []int{2, 3, 1, 1, 4},
			want: 2,
		},
		{
			in:   []int{2, 3, 0, 1, 4},
			want: 2,
		},
		{
			in:   []int{2, 4, 3, 1, 1, 2, 9, 1, 7},
			want: 4,
		},
		{
			in:   []int{3, 5, 4, 3, 2, 1, 1, 5, 9},
			want: 4,
		},
	}
	for _, tc := range tcases {
		if got := jump2(tc.in); tc.want != got {
			t.Logf("Test case %v failed, want : %d, got %d\n", tc, tc.want, got)
			t.Fail()
		}
	}
}

/*
*
The below is the more optimal soln
jump to the next idx which gives u more net leaverage or power
power = maxof(idx + nums[idx])
if we have more than one idx have same net effective power, always choose the one farthest from current index
*/
func jump2(nums []int) int {
	size := len(nums)
	if size <= 1 {
		return 0
	}
	nextIdx := func(cIdx int) int {
		num := nums[cIdx]
		if num+cIdx >= size-1 {
			return size - 1
		}
		netBigIdx := 0
		for netBig, nIdx, untill := 0, cIdx+1, cIdx+num; nIdx <= untill; nIdx++ {
			if tmp := nums[nIdx] + nIdx; tmp >= netBig {
				netBig = tmp
				netBigIdx = nIdx
			}
		}
		return netBigIdx
	}
	for nIdx, jumps, lIdx := 0, 0, size-1; ; {
		nIdx = nextIdx(nIdx)
		jumps++
		if nIdx >= lIdx {
			return jumps
		}
	}
}

func TestChannelAsQueue(t *testing.T) {
	var queue chan int = make(chan int, 11)
	for i := 0; i < 10; i++ {
		queue <- i
	}
	for i := 0; i < 10; i++ {
		t.Log(<-queue)
	}
}

func TestMemoMap(t *testing.T) {
	var memo map[int]bool = make(map[int]bool)
	memo[1] = true
	memo[2] = false
	key := 2
	if flag, ok := memo[key]; ok {
		t.Logf("%v exists in memo map, its value %v\n", key, flag)
	} else {
		t.Log("3 does not exists in map")
	}
}
