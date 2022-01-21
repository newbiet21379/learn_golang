package leetcode

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	var (
		res        [][]int
		i, j       int
		loIn, hiIn int
	)
	for i < len(firstList) && j < len(secondList) {
		// Check intersect
		//Low Intersect
		if firstList[i][0] > secondList[j][0] {
			loIn = firstList[i][0]
		} else {
			loIn = secondList[j][0]
		}
		// High Intersect
		if firstList[i][1] < secondList[j][1] {
			hiIn = firstList[i][1]
		} else {
			hiIn = secondList[j][1]
		}
		if loIn <= hiIn {
			res = append(res, []int{loIn, hiIn})
		}
		if firstList[i][1] < secondList[j][1] {
			i++
		} else {
			j++
		}
	}
	return res
}
