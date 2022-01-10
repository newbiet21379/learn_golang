package leetcode

func SortSentence(s string) string {
	var (
		n    = len(s)
		pos  int
		word []byte
	)
	words := make(map[int]string)
	for ; pos < n; pos++ {
		if isDigit(s[pos]) {
			digit := int(s[pos] - '0')
			words[digit] = string(word)
			word = nil
			continue
		}
		if s[pos] != ' ' {
			word = append(word, s[pos])
		}
	}
	for i := 1; i <= len(words); i++ {
		word = append(word, words[i]...)
		if i < len(words) {
			word = append(word, ' ')
		}
	}
	return string(word)
}
