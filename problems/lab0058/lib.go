package lab0058

func LengthOfLastWord(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 32 {
			if len(s)-i-1 == 0 {
				s = s[:len(s)-1]
				continue
			} else {
				return len(s) - i - 1
			}
		}
	}

	return len(s)
}
