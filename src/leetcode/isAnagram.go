package leetcode

import (
	"sort"
	"strings"
)

func isAnagram2(s string, t string) bool {
	charList := make([]int, 26)
	for i := 0; i < len(s); i++ {
		charList[s[i]-'a']++
	}
	for i := 0; i < len(t); i++ {
		charList[t[i]-'a']--
	}
	for _, num := range charList {
		if num != 0 {
			return false
		}
	}
	return true
}
func IsAnagram(s string, t string) bool {
	cache := make(map[byte]int)
	if len(s) != len(t) {
		return false
	}
	for i := range s {
		_, ok := cache[s[i]]
		if ok {
			cache[s[i]]++
			continue
		}
		cache[s[i]] = 1
	}
	for i := range t {
		_, ok := cache[t[i]]
		if ok {
			cache[t[i]]--
		}
	}
	for _, v := range cache {
		if v != 0 {
			return false
		}
	}
	return true
}

func Len(s string) int {
	return len(s)
}

type Str []byte

func (a Str) Len() int           { return len(a) }
func (a Str) Less(i, j int) bool { return a[i] < a[j] }
func (a Str) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func IsAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	a := append(Str{}, s...)
	b := append(Str{}, t...)
	sort.Sort(a)
	sort.Sort(b)
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var arr [256]rune
	for i := range s {
		arr[s[i]]++
		arr[t[i]]--
	}
	for i := range arr {
		if arr[i] != 0 {
			return false
		}
	}
	return true
}

func groupAnagrams(strs []string) [][]string {
	var (
		res   [][]string
		cache = make(map[string][]string)
		sortS []byte
	)
	for _, s := range strs {
		sortS = []byte(s)
		sort.Slice(sortS, func(i, j int) bool {
			return sortS[i] < sortS[j]
		})
		_, ok := cache[string(sortS)]
		if ok {
			cache[string(sortS)] = append(cache[string(sortS)], s)
		} else {
			cache[string(sortS)] = []string{s}
		}
	}
	for _, v := range cache {
		sort.Slice(v, func(i, j int) bool {
			return strings.Compare(v[i], v[j]) < 0
		})
		res = append(res, v)
	}
	sort.Slice(res, func(i, j int) bool {
		return len(res[i]) < len(res[j])
	})
	return res
}
