package lab0066

func PlusOne(digits []int) []int {
	carry := false
	for i := len(digits) - 1; i >= 0; i-- {
		if i == len(digits)-1 {
			digits[i] += 1
		} else {
			if carry {
				digits[i] += 1
			}
		}

		if digits[i] == 10 {
			digits[i] = 0
			if i == 0 {
				return append([]int{1}, digits...)
			}
			carry = true
		} else {
			carry = false
		}

		if !carry {
			break
		}
	}

	return digits
}
