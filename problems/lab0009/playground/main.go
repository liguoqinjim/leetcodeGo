package main

import (
	"fmt"
	"strconv"
)

func main() {
	//fmt.Println(isPalindrome(1221))

	fmt.Println(isPalindrome2(21112))
}

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func isPalindrome2(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}
	rev := 0
	for x > rev {
		rev = rev*10 + x%10
		x = x / 10
	}
	fmt.Println("rev=", rev, "x=", x)
	return (x == rev || x == rev/10)
}
