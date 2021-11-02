package main

import (
	"fmt"
)
func findShortestCombination( n int , k int) int {
	var listTest []int
	count := 0
	temp := k
	if n == k {
		return 1
	}

	for index := n; index >= 1; index-- {
		if temp >= index {
			temp = temp - index
			count++
			listTest = append(listTest, index)
		} else {
			continue
		}
	}

	if temp > 0 {
		return -1
	} else {
		return count
	}
}

func ShortestSubString(input string) int{
	var allSubString []string
	for i := 0;i < len(input);i++{
		for j := i+1;j<=len(input);j++{
			allSubString = append(allSubString,input[i:j])
		}
	}
	var mapSubStringLen map[string] int = make(map[string] int)
	for _,item := range allSubString {
		_, ok := mapSubStringLen[item]
		if ok == true{
			mapSubStringLen[item] ++
		}else{
			mapSubStringLen[item] = 1
		}
	}
	minOccurance := len(input)
	var minText string = input

	for key, value := range mapSubStringLen{
		if value <= minOccurance{
			minOccurance = value
			if len(key) <= len(minText) && minOccurance == 1{
				minText = key
			}
		}
	}
	return len(minText)
}

func main() {
	fmt.Println(findShortestCombination(5,10))
	fmt.Println(findShortestCombination(10,55))
	fmt.Println(findShortestCombination(5,15))
	fmt.Println(ShortestSubString("aabbbabaaa"))
}