package leetcode

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

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

func MiniMaxSum(arr []int32) {
	// Write your code here
	var (
		sum int64
		max int64 = 0
		min int64 = 1 << 32
	)
	for _, v := range arr {
		sum += int64(v)
	}
	for _, v := range arr {
		if sum > max+int64(v) {
			max = sum - int64(v)
		}
		if sum < min+int64(v) {
			min = sum - int64(v)
		}
	}
	fmt.Printf("%v %v", min, max)
}

func birthdayCakeCandles(candles []int32) int32 {
	// Write your code here
	var (
		max, count int32
	)
	for _, v := range candles {
		if v > max {
			max = v
		}
	}
	for _, v := range candles {
		if v == max {
			count++
		}
	}
	return count
}

func TimeConversion(s string) string {
	// Write your code here
	noon, hour := strings.HasSuffix(s, "PM"), s[:2]
	h, _ := strconv.Atoi(hour)
	if h == 12 && !noon {
		h -= 12
	} else if h != 12 && noon {
		h += 12
	}
	if h < 10 {
		hour = string([]byte{'0', byte(h + '0')})
	} else {
		hour = strconv.Itoa(h)
	}
	res := strings.Replace(s, s[:2], hour, 1)
	return res[:len(res)-2]
}
