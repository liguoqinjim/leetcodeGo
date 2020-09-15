package lab0068

func MySqrt(x int) int {
	for i := (x + 2) / 2; i > 0; i-- {
		if (i * i) < 0 {
			continue
		}

		if i*i > x && (i-1)*(i-1) < x {
			return i - 1
		} else if i*i == x {
			return i
		} else if (i-1)*(i-1) == x {
			return i - 1
		}
	}

	return 0
}
