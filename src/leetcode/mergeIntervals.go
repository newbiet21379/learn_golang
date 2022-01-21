package leetcode

import "sort"

func Merge(intervals [][]int) [][]int {
	var merged [][]int
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for _, v := range intervals {
		if len(merged) == 0 || merged[len(merged)-1][1] < v[0] {
			merged = append(merged, v)
		} else {
			if merged[len(merged)-1][1] < v[1] {
				merged[len(merged)-1][1] = v[1]
			}
		}

	}
	return merged
}
