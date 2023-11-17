package dynamicprogram

import (
	"math"
	"testing"
)

/*
 * https://leetcode.com/problems/interleaving-string/
 * Given strings s1, s2 and s3, find whether s3 is formed by an interleaving of s1 and s2.
 * An interleaving of two strings s and t is a configuration where s and t are divided into n amd m substrings respectivel, such that
 * s = s1 + s2 + ... + sn
 * t = t1 + t2 + ... + tn
 * |n - m| <= 1
 * the interleaving is s1 + t1 + s2 + t2 + s3 + t3 + ... or t1 + s1 + t2 + s2 + t3 + s3 + ...const
 *
 * Note:  a + b is the concatenation of string a and b
 *
 * Constraints:0 <= s1.length, s2.length <= 100
 * 0 <= s3.length <= 200
 * s1, s2, and s3 consist of lowercase English letters.
 *
 */
func isInterleave(s1 string, s2 string, s3 string) bool {
	l1, l2, l3 := len(s1), len(s2), len(s3)
	if l3 != l1+l2 {
		return false
	}
	p1, p2, p3 := 0, 0, 0
	n, m := 0, 0
	for {
		if p3 == l3 {
			break
		}
		if p1 == l1 && p2 == l2 {
			break
		}
		if p1 < l1 && s1[p1] == s3[p3] {
			for p1 < l1 && p3 < l3 && s1[p1] == s3[p3] {
				p1++
				p3++
			}
			n++
		} else if p2 < l2 && s2[p2] == s3[p3] {
			for p2 < l2 && p3 < l3 && s2[p2] == s3[p3] {
				p2++
				p3++
			}
			m++
		} else {
			return false
		}
	}
	if p1 == l1 && p2 == l2 && p3 == l3 && math.Abs(float64(n-m)) <= 1 {
		return true
	} else {
		return false
	}
}

func TestIsInterleave(t *testing.T) {
	tcases := []struct {
		s1   string
		s2   string
		s3   string
		want bool
	}{
		{
			s1:   "aabcc",
			s2:   "dbbca",
			s3:   "aadbbcbcac",
			want: true,
		},
		{
			s1:   "aabcc",
			s2:   "dbbca",
			s3:   "aadbbbaccc",
			want: false,
		},
		{
			s1:   "",
			s2:   "",
			s3:   "",
			want: true,
		},
		{
			s1:   "",
			s2:   "b",
			s3:   "b",
			want: true,
		},
		{
			s1:   "a",
			s2:   "",
			s3:   "a",
			want: true,
		},
		{
			s1:   "aa",
			s2:   "ab",
			s3:   "aaba",
			want: true,
		},
	}
	for _, tc := range tcases {
		if got := isInterleave(tc.s1, tc.s2, tc.s3); got != tc.want {
			t.Logf("Testcase %v failed, got: %v\n", tc, got)
			t.Fail()
		}
	}
}
