package leetcode

func FindMinNormal(nums []int) int {
	var (
		low, high = 0, len(nums) - 1
	)
	if len(nums) == 1 {
		return nums[0]
	}
	for low < high {
		mid := low + (high-low)/2

		//if nums[low] == nums[high] {
		//	low++
		//} else if nums[mid] > nums[high] {
		//	low = mid + 1
		//} else {
		//	high = mid
		//}
		if nums[mid] < nums[high] {
			high = mid
		} else if nums[mid] > nums[high] {
			low = mid + 1
		}
	}
	return nums[low]
}

func FindMinHard(nums []int) int {
	var (
		low, high = 0, len(nums) - 1
	)
	if len(nums) == 1 {
		return nums[0]
	}
	for low < high {
		mid := low + (high-low)/2

		//if nums[low] == nums[high] {
		//	low++
		//} else if nums[mid] > nums[high] {
		//	low = mid + 1
		//} else {
		//	high = mid
		//}
		if nums[mid] < nums[high] {
			high = mid
		} else if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high--
		}
	}
	return nums[low]
}
