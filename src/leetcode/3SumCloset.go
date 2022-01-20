package leetcode

//
//func threeSumClosest(nums []int, target int) int {
//	var (
//		result int
//		n      = len(nums)
//		diff   = math.MaxUint
//	)
//	sort.Ints(nums)
//	for i := range nums {
//		leftI, rightI := i+1, n-1
//		if i > 0 && nums[i] == nums[i-1] {
//			continue
//		}
//		for rightI > leftI {
//			sum := nums[i] + nums[leftI] + nums[rightI]
//			if sum < target {
//				leftI++
//			} else {
//				rightI--
//			}
//			if int(math.Abs(float64(sum-target))) < diff {
//				result = sum
//				diff = int(math.Abs(float64(sum - target)))
//			}
//		}
//	}
//	return result
//}
