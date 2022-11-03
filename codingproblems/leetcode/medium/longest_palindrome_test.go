package medium

import (
	"testing"
)

/*
https://leetcode.com/problems/longest-palindrome-by-concatenating-two-letter-words/
You are given an array of strings words. Each element of words consists of two lowercase English letters.
Create the longest possible palindrome by selecting some elements from words and concatenating them in any order. Each element can be selected at most once.
Return the length of the longest palindrome that you can create. If it is impossible to create any palindrome, return 0.

A palindrome is a string that reads the same forward and backward.

Constraints:
============
1 <= words.length <= 105
words[i].length == 2
words[i] consists of lowercase English letters.
*/

func longestpalindrome(words []string) int {
	var length int = 0
	var wordmap map[string]int = map[string]int{}

	for _, word := range words {
		rev := reverse(word)
		if count, exists := wordmap[rev]; exists {
			length = length + 4
			if count == 1 {
				delete(wordmap, rev)
			} else {
				wordmap[rev] = count - 1
			}
		} else {
			wordmap[word] = wordmap[word] + 1
		}
	}

	for word, _ := range wordmap {
		if word[0] == word[1] {
			length = length + 2
			break
		}
	}
	return length
}

func reverse(word string) string {
	return string([]byte{word[1], word[0]})
}

func TestLongestPalindrome(t *testing.T) {
	tcases := []struct {
		in   []string
		want int
	}{
		{
			in:   []string{"ab"},
			want: 0,
		},
		{
			in:   []string{"cc"},
			want: 2,
		},
		{
			in:   []string{"ab", "df", "ba"},
			want: 4,
		},
		{
			in:   []string{"ab", "cc", "ba"},
			want: 6,
		},
		{
			in:   []string{"ab", "ab", "ba", "ba", "ab", "cc"},
			want: 10,
		},
		{
			in:   []string{"dd", "aa", "bb", "dd", "aa", "dd", "bb", "dd", "aa", "cc", "bb", "cc", "dd", "cc"},
			want: 22,
		},
	}
	for _, tcase := range tcases {
		got := longestpalindrome(tcase.in)
		if got != tcase.want {
			t.Errorf("TC Failed. input: %v, got: %d, want:%d\n", tcase.in, got, tcase.want)
		}
	}
}

func BenchmarkLongestPalindrome(b *testing.B) {
	var in []string = []string{"dd", "aa", "bb", "dd", "aa", "dd", "bb", "dd", "aa", "cc", "bb", "cc", "dd", "cc"}
	var got int = 0
	for i := 0; i < b.N; i++ {
		got = longestpalindrome(in)
	}
	got++
}
