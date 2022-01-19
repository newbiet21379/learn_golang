package leetcode

import "sort"

//Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
//
//Notice that the solution set must not contain duplicate triplets.
//
//
//
//Example 1:
//
//Input: nums = [-1,0,1,2,-1,-4]
//Output: [[-1,-1,2],[-1,0,1]]
//Example 2:
//
//Input: nums = []
//Output: []
//Example 3:
//
//Input: nums = [0]
//Output: []
//
//
//Constraints:
//
//0 <= nums.length <= 3000
//-105 <= nums[i] <= 105
// Place 3 pointer started with i , i+1 and len - 1
// for same value skip next
// in loop check pointer for closet sum 
func threeSum(nums []int) [][]int {
	var (
		triplets [][]int
		n        = len(nums)
	)
	if n == 0 || n == 1 {
		return triplets
	}
	sort.Ints(nums)
	for i := range nums {
		leftI, rightI := i+1, n-1
		if i > 0 && nums[i] == nums[i-1] {
			continue
		} else {
			for rightI > leftI {
				sum := nums[i] + nums[leftI] + nums[rightI]
				if sum > 0 {
					rightI--
				} else if sum < 0 {
					leftI++
				} else {
					triplets = append(triplets, []int{nums[i], nums[leftI], nums[rightI]})
					leftI++
					for nums[leftI] == nums[leftI-1] && leftI < rightI {
						leftI++
					}
				}
			}
		}
	}
	return triplets
}
