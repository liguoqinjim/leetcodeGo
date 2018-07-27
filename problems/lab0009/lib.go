package lab0009

import (
	"strconv"
)

//这个解不好，因为题目里面说最好能不把数字转为字符串处理
func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

//sample
func isPalindrome2(x int) bool {
	if x < 0 || x%10 == 0 {
		return false
	}
	res := 0
	for x > res {
		res = res*10 + x%10
		x /= 10
	}
	return x == res || res/10 == x
}
