package leetcode

import "strconv"

func IsPalindrome(x int) bool {

	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	isPalindrome := true
	xStr := strconv.Itoa(x)
	xStrLen := len(xStr)

	for i := 0; i < xStrLen; i++ {
		if xStr[i] != xStr[xStrLen-1-i] {
			isPalindrome = false
			break
		}
	}

	return isPalindrome
}
