package leetcode

func CanPlaceFlowers(flowerbed []int, n int) bool {
	var (
		count, i int
		l        = len(flowerbed)
	)

	for i < l {
		if flowerbed[i] == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == l-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			i++
			count++
		}
		if count >= n {
			return true
		}
		i++
	}
	return false
}
