package lab0035

func searchInsert(nums []int, target int) int {
	for n, v := range nums {
		if target <= v {
			return n
		}
	}
	return len(nums)
}

//sample

//时间复杂度O(logn)
func searchInsert2(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}
