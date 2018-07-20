package lab0007

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

//sample
//优化，不用另一个slice参与计算，overflow的处理也更简单
func reverse2(x int) int {
	var result int32 = 0
	for x != 0 {
		tail := int32(x) % 10
		newResult := result*10 + tail
		if (newResult-tail)/10 != result { //overflow，因为从数学来考虑，这个等式是肯定相等的。如果不相等，就出现了数字超过范围的情况了
			return 0
		}
		result = newResult
		x /= 10
	}
	return int(result)
}
