package lab007

import (
	"math"
)

func reverse(x int) int {
	if x >= math.MaxInt32 || x <= math.MinInt32 {
		return 0
	}

	i := 10
	num := make([]int, 0)
	for {
		rem := x % i
		num = append(num, rem/(i/10))
		if rem == x {
			break
		}
		i *= 10
	}
	result := 0
	for n, v := range num {
		result += v * int(math.Pow(10, float64(len(num)-1-n)))
	}

	if result >= math.MaxInt32 || result <= math.MinInt32 {
		return 0
	}

	return result
}
