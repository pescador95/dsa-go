package leetcode

func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}

	low, maxLen := 0, 0

	expand := func(left, right int) {
		for left >= 0 && right < n && s[left] == s[right] {
			left--
			right++
		}

		if right-left-1 > maxLen {
			low = left + 1
			maxLen = right - left - 1
		}
	}

	for i := 0; i < n-1; i++ {
		expand(i, i)
		expand(i, i+1)
	}

	return s[low : low+maxLen]
}

// Given a string s, return the longest palindromic substring in s.

// Example 1:

// Input: s = "babad"
// Output: "bab"
// Explanation: "aba" is also a valid answer.
// Example 2:

// Input: s = "cbbd"
// Output: "bb"

// Constraints:

// 1 <= s.length <= 1000
// s consist of only digits and English letters.
