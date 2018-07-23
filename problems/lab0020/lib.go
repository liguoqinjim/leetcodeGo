package lab0020

func isValid(s string) bool { //这个解法用到了acsii值
	if len(s) == 0 {
		return true
	}

	left := make([]int, 1)
	left[0] = int(s[0])
	for i := 1; i < len(s); i++ {
		a := int(s[i])
		if len(left) == 0 {
			left = append(left, a)
			continue
		}
		if a-left[len(left)-1] != 1 && a-left[len(left)-1] != 2 {
			left = append(left, a)
		} else {
			left = left[:len(left)-1]
		}
	}

	if len(left) == 0 {
		return true
	} else {
		return false
	}
}
