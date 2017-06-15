package lib

//时间复杂度是O(n)
func twoSum2(nums []int, target int) []int {
	r := make([]int, 2)
	result := make(map[int]int) //key是nums的值,value是值对应的下标
	for n, v := range nums {
		if value, ok := result[target-v]; ok {
			r[0] = value
			r[1] = n
			return r
		}
		result[v] = n
	}

	return r
}
