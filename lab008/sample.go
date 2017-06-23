package lab008

import (
	"math"
)

func myAtoi2(str string) int {
	//empty string
	if len(str) == 0 {
		return 0
	}

	index, sign, result := 0, 1, 0
	//remove space
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			index++
		} else {
			break
		}
	}

	//handle sign
	if str[index] == '+' || str[index] == '-' {
		if str[index] == '-' {
			sign = -1
		}
		index++
	}

	//Convert number and avoid overflow
	for i := index; i < len(str); i++ {
		digit := int(str[i] - '0')
		if digit < 0 || digit > 9 {
			break
		}

		if sign == 1 {
			if math.MaxInt32/10 == result || math.MaxInt32/10 < result && math.MaxInt32%10 < digit {
				return math.MaxInt32
			}
		} else {
			if (math.MaxInt32/10 == result || math.MaxInt32/10 < result) && digit > 8 {
				return math.MinInt32
			}
		}

		result = result*10 + digit
		index++
	}

	return result * sign
}
