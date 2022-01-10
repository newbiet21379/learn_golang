package leetcode

func AreNumbersAscending(s string) bool {
	var (
		currentNumber, last, pos int
		n                        = len(s)
	)
	for pos < n {
		if isDigit(s[pos]) {
			digit := int(s[pos] - '0')
			currentNumber = digit
			if pos+1 < n && isDigit(s[pos+1]) {
				digit = int(s[pos+1] - '0')
				currentNumber = currentNumber*10 + digit
				pos++
			}
			if currentNumber <= last {
				return false
			}
			last = currentNumber
		}
		pos++
	}
	return true
}
