package dynamicprogram

import (
	"testing"
)

/**
 * https://leetcode.com/problems/longest-palindromic-substring/
 * Given a string s, return the longest palindromic substring in s.
 * Constraints:
 * 1 <= s.length <= 1000
 * s consist of only digits and English letters.
 */
func longestPalindrome(s string) string {
	for size, inc := len(s), 0; size > 1; size-- {
		for j := 0; j <= inc; j++ {
			if key := s[j : size+j]; isPalindrome(key) {
				return key
			}
		}
		inc++
	}
	return string(s[0])
}

func isPalindrome(s string) bool {
	if size := len(s); size <= 1 {
		return false
	} else {
		for start, end := 0, size-1; start < end; {
			if s[start] != s[end] {
				return false
			}
			start++
			end--
		}
	}
	return true
}

func TestLongestPalindrome(t *testing.T) {
	tcases := []struct {
		input    string
		expected string
	}{
		{
			input:    "ana",
			expected: "ana",
		},
		{
			input:    "banana",
			expected: "anana",
		},
		{
			input:    "haah",
			expected: "haah",
		},
		{
			input:    "vimal",
			expected: "v",
		},
		{
			input:    "aaa",
			expected: "aaa",
		},
		{
			input:    "babad",
			expected: "bab",
		},
	}
	for _, tc := range tcases {
		if actual := longestPalindrome(tc.input); tc.expected != actual {
			t.Logf("test case %v failed, expected: %s got: %s\n", tc, actual, tc.expected)
			t.Fail()
		}
	}
}

func TestPalindrome(t *testing.T) {
	longestPalindrome("banana")
}
