package leetcode

func isValid(s string) bool {
	code := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	var left []byte
	for i := 0; i < len(s); i++ {
		if _, ok := code[s[i]]; ok {
			left = append(left, s[i])
		} else {
			if len(left) == 0 {
				return false
			}
			lend := left[len(left)-1]
			reverse, _ := code[lend]
			if reverse == s[i] {
				left = left[:len(left)-1]
			} else {
				break
			}
		}
	}
	return len(left) == 0
}
