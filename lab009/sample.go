package lab009

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
