package lab007

//优化，不用另一个slice参与计算，overflow的处理也更简单
func reverse2(x int) int {
	var result int32 = 0
	for x != 0 {
		var tail int32 = int32(x) % 10
		var newResult int32 = result*10 + tail
		if (newResult-tail)/10 != result { //overflow
			return 0
		}
		result = newResult
		x /= 10
	}
	return int(result)
}
