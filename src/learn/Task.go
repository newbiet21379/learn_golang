package learn

type Task struct {
	Capacity int
	Total int
}

func (t Task) FindShortestCombination(task Task) int {
	count := 0
	temp := task.Total
	if task.Capacity == task.Total {
		return 1
	}

	for index := task.Capacity; index >= 1; index-- {
		if temp >= index {
			temp = temp - index
			count++
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

func (t Task) ShortestSubString(input string) int{
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

