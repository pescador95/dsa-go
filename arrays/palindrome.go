package main

import (
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	low := 0
	high := len(s) - 1

	for low < high {
		if s[low] != s[high] {
			return false
		}
		low++
		high--
	}
	return true
}
