package lab0038

import "strconv"

func CountAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	if n == 2 {
		return "11"
	}

	last := CountAndSay(n - 1)
	result := ""
	num := 0
	nowValue := '0'
	for n, v := range last {
		if n == 0 {
			nowValue = v
			num = 1
			continue
		}

		if v != nowValue {
			result += strconv.Itoa(num) + string(nowValue)

			nowValue = v
			num = 1
		} else {
			num += 1
		}

		if n == len(last)-1 {
			result += strconv.Itoa(num) + string(nowValue)
		}
	}

	return result
}
