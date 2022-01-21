package leetcode

import "math"

func maxProfit(prices []int) int {
	var (
		profit int
		min    = math.MaxInt
	)
	for i := range prices {
		if prices[i] < min {
			min = prices[i]
		} else {
			if prices[i]-min > profit {
				profit = prices[i] - min
			}
		}
	}
	return profit
}
