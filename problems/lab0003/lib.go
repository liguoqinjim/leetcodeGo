package lab0003

func LengthOfLongestSubstring(s string) int {
	max := 0
	l := 0

	lastStartIndex := 0
	m := make(map[uint8]int)

	for i := 0; i < len(s); {
		if len(m) == 0 {
			lastStartIndex = i
		}

		v := s[i]

		if _, ok := m[v]; ok {
			i = lastStartIndex + 1
			m = make(map[uint8]int)
			if l > max {
				max = l
			}
			l = 0
			continue
		} else {
			l += 1
			m[v] = 1
		}

		if i == len(s)-1 {
			if l > max {
				max = l
			}
		}

		i++
	}

	return max
}

func LengthOfLongestSubstring2(s string) int {
	max := 0

	start := 0
	m := make(map[uint8]int)

	for i := 0; i < len(s); i++ {
		v := s[i]

		if index, ok := m[v]; ok && index >= start {
			start = index + 1
		}
		if max < (i - start + 1) {
			max = i - start + 1
		}

		m[v] = i
	}

	return max
}
