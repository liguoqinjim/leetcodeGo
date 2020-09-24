package lab0068

import "math"

//会超时
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

//袖珍计算器算法
func MySqrt2(x int) int {
	if x == 0 {
		return 0
	}
	ans := int(math.Exp(0.5 * math.Log(float64(x))))
	if (ans+1)*(ans+1) <= x {
		return ans + 1
	}
	return ans
}

//二分法
func MySqrt3(x int) int {
	l, r := 0, x
	ans := -1
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}
