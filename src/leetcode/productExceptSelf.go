package leetcode

func ProductExceptSelf(nums []int) []int {
	var (
		output        = make([]int, len(nums))
		product, zero = 1, -1
	)
	for i := range nums {
		if nums[i] == 0 {
			if zero > -1 {
				return output
			}
			zero = i
		} else {
			product *= nums[i]
		}
	}

	if zero > -1 {
		output[zero] = product
		return output
	}

	for i := range nums {
		output[i] = product / nums[i]
	}

	return output
}
