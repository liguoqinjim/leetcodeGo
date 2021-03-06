package lab0027

func RemoveElement(nums []int, val int) int {
	l := 0
	i := 0

	for i != len(nums) {
		if nums[i] != val {
			l++
			i++
		} else {
			if i == len(nums)-1 {
				nums = nums[:i]
			} else {
				nums[i] = nums[len(nums)-1]
				nums = nums[:len(nums)-1]
			}
		}
	}

	return l
}

func RemoveElement2(nums []int, val int) int {
	l := 0

	for _, v := range nums {
		if v != val {
			nums[l] = v
			l++
		}
	}

	return l
}
