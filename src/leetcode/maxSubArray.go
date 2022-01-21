package leetcode

func maxSubArray(nums []int) int {
	var (
		sum    int
		maxSum = -1 << 10
	)
	for i := range nums {
		sum += nums[i]
		if sum < 0 {
			sum = 0
		}
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}
