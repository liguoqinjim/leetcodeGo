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

func isValid2(s string) bool {
	var stack []byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch c {
		case '(', '[', '{':
			stack = append(stack, c)
		case ')', ']', '}':
			l := len(stack)
			if l == 0 {
				return false
			}
			v := stack[l-1]
			stack = stack[:l-1]
			if c == ')' && v != '(' {
				return false
			}

			if c == ']' && v != '[' {
				return false
			}

			if c == '}' && v != '{' {
				return false
			}
		}
	}

	if len(stack) > 0 {
		return false
	}

	return true
}
