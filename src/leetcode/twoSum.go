package leetcode

func TwoSum(nums []int, target int) []int {
	var n = len(nums)

	cache := make(map[int]int, n)

	for i := range nums {
		a := nums[i]
		j, ok := cache[target-a]
		if ok {
			return []int{i, j}
		}
		cache[a] = i
	}
	panic("No solution found")
}

func TwoSum2(nums []int, target int) []int {
	var (
		i, j int
		res  []int
		n    = len(nums)
	)
	for i < n {
		left := target - nums[i]
		for j < n && left != nums[j] {
			j++
		}
		if j < n && nums[j] == left && j != i {
			if i > j {
				res = append(res, j, i)
			} else {
				res = append(res, i, j)
			}
			break
		} else {
			i++
			j = 0
		}
	}
	return res
}
