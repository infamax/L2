package sorts

import (
	"sort"
	"strings"
)

const countMonth = 12

var monthsByName = map[string]int{
	"january":   0,
	"february":  1,
	"march":     2,
	"april":     3,
	"may":       4,
	"june":      5,
	"july":      6,
	"august":    7,
	"september": 8,
	"october":   9,
	"november":  10,
	"december":  11,
}

var monthByHisNumber = map[int]string{
	0:  "january",
	1:  "february",
	2:  "march",
	3:  "april",
	4:  "may",
	5:  "june",
	6:  "july",
	7:  "august",
	8:  "september",
	9:  "october",
	10: "november",
	11: "december",
}

func SortMonth(input []string) []string {
	monthSlice, otherItemSlice := divideByTwoArrays(input)
	sort.Strings(otherItemSlice)
	countSort(monthSlice)
	return mergeSlices(otherItemSlice, monthSlice)
}

func countSort(monthSlice []string) {
	var arr [countMonth]int
	for _, month := range monthSlice {
		arr[monthsByName[month]]++
	}

	k := 0
	for i := 0; i < countMonth; i++ {
		for arr[i] > 0 {
			monthSlice[k] = monthByHisNumber[i]
			arr[i]--
			k++
		}
	}
}

func mergeSlices(arr1, arr2 []string) []string {
	res := make([]string, 0, len(arr1)+len(arr2))
	for i := 0; i < len(arr1); i++ {
		res = append(res, arr1[i])
	}
	for i := 0; i < len(arr2); i++ {
		res = append(res, arr2[i])
	}
	return res
}

func divideByTwoArrays(input []string) ([]string, []string) {
	monthSlice := make([]string, 0)
	otherItemSlice := make([]string, 0)
	for _, str := range input {
		_, ok := monthsByName[strings.ToLower(str)]
		if ok {
			monthSlice = append(monthSlice, str)
		} else {
			otherItemSlice = append(otherItemSlice, str)
		}
	}
	return monthSlice, otherItemSlice
}
