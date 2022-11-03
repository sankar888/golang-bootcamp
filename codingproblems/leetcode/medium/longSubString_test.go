package medium

import (
	"testing"
)

/*
https://leetcode.com/problems/longest-substring-without-repeating-characters/
Given a string s, find the length of the longest substring without repeating characters.
*/
func lengthOfLongestSubstring(s string) int {
	var llen int = 0
	var length int = len(s)
	for i := 0; i < length; i++ {
		j := i                                   //new map
		var cmap map[byte]bool = map[byte]bool{} //only ascii strings
		for ; j < length; j++ {
			if cmap[s[j]] {
				sstrlen := j - i
				if sstrlen > llen {
					llen = sstrlen
				}
				break
			}
			cmap[s[j]] = true
		}
		//if end is reached in one of the iteration, no need to continue
		if j == length {
			sstrlen := j - i
			if sstrlen > llen {
				llen = sstrlen
			}
			break
		}
	}
	return llen
}

func lengthOfLongestSubstringSol1(s string) int {
	if len(s) == 0 {
		return 0
	}
	var b [256]bool
	res, l, r := 0, 0, 0
	for l+res < len(s) {
		if b[s[r]] {
			b[s[l]] = false
			l++
		} else {
			b[s[r]] = true
			r++
		}
		if res < r-l {
			res = r - l
		}
		if r >= len(s) {
			break
		}
	}
	return res
}

var tcases = []struct {
	in  string
	out int
}{
	{"abc", 3},
	{"aabca", 3},
	{"bbbb", 1},
	{"abadefg", 6},
}

func TestLengthOfLongestSubstring(t *testing.T) {
	for _, tcase := range tcases {
		out := lengthOfLongestSubstring(tcase.in)
		if tcase.out == out {
			t.Logf("TC passed: inp: %s, expected: %d, actual : %d \n", tcase.in, tcase.out, out)
		} else {
			t.Errorf("TC Failed: inp: %s, expected: %d, actual : %d \n", tcase.in, tcase.out, out)
		}

	}
}

func TestLengthOfLongestSubstringSol1(t *testing.T) {
	for _, tcase := range tcases {
		out := lengthOfLongestSubstringSol1(tcase.in)
		if tcase.out == out {
			t.Logf("TC passed: inp: %s, expected: %d, actual : %d \n", tcase.in, tcase.out, out)
		} else {
			t.Errorf("TC Failed: inp: %s, expected: %d, actual : %d \n", tcase.in, tcase.out, out)
		}

	}
}

func BenchmarkLengthOfLongestSubstringSol1(b *testing.B) {
	var s string = "abadefg"
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstringSol1(s)
	}
}

func BenchmarkLengthOfLongestSubstring(b *testing.B) {
	var s string = "abadefg"
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring(s)
	}
}

/*
TODO: why is my method slow ? What is the logic behing the fastest soln
TODO: what i need to learn / practice
BenchmarkLengthOfLongestSubstringSol1-8   	158421430	         7.803 ns/op	       0 B/op	       0 allocs/op
BenchmarkLengthOfLongestSubstring-8       	 5809292	       211.5 ns/op	       0 B/op	       0 allocs/op
*/
