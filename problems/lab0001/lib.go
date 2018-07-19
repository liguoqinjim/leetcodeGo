package lib

func twoSum(nums []int, target int) []int {
	as := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				as[0] = i
				as[1] = j
			}
		}
	}

	return as
}

//时间复杂度是O(n)
func twoSum2(nums []int, target int) []int {
	r := make([]int, 2)
	result := make(map[int]int) //key是nums的值,value是值对应的下标
	for n, v := range nums {
		if value, ok := result[target-v]; ok {
			r[0] = value
			r[1] = n
			break
		}
		result[v] = n
	}

	return r
}
