package leetcode

func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1 / myPow(x, -n)
	}

	res := float64(1)

	for n > 0 {
		if n%2 == 0 {
			x = x * x
			n = n / 2
		} else {
			res = res * x
			n--
		}
	}

	return res
}
