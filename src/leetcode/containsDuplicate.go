package leetcode

func containsDuplicate(nums []int) bool {
	cache := make(map[int]int)
	for i := range nums {
		a := nums[i]
		_, ok := cache[a]
		if ok {
			return true
		}
		cache[a] = i
	}
	return false
}
