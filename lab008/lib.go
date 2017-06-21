package lab008

import "math"

func myAtoi(str string) int {
	if len(str) == 0 {
		return 0
	}
	result := 0
	u := -1 //有无符号
	e := -1 //错误位第几位
	for n, v := range str {
		if v < 48 || v > 57 {
			if v == 43 || v == 45 { //+-号
				if result != 0 {
					e = n
					break
				}
				if u >= 0 {
					return 0
				}
				u = n
				continue
			}
			if v == 32 { //空格
				if result != 0 {
					e = n
					break
				}
				if u >= 0 { //这里默认符号之后只能跟数字
					return 0
				}
				continue
			}
			e = n
			break
		}
		result += int(v-48) * int(math.Pow(10, float64(len(str)-n-1)))
	}
	if e >= 0 {
		result /= int(math.Pow(10, float64(len(str)-e)))
	}
	if u >= 0 {
		if str[u] == 45 {
			result = -result
		}
	}
	if result > math.MaxInt32 {
		if u >= 0 {
			if str[u] == 45 {
				return math.MinInt32
			}
		}
		result = math.MaxInt32
	}
	if result < math.MinInt32 {
		if u < 0 {
			return math.MaxInt32
		} else {
			if str[u] != 45 {
				return math.MaxInt32
			}
		}
		result = math.MinInt32
	}

	return result
}
