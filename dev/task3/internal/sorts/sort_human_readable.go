package sorts

import (
	"sort"
	"strconv"
)

var m = map[byte]int64{
	'B': 8,
	'K': 1024,
	'M': 1024 * 1024,
	'G': 1024 * 1024 * 1024,
	'T': 1024 * 1024 * 1024 * 1024,
	'P': 1024 * 1024 * 1024 * 1024 * 1024,
}

func SortHumanReadable(input []string) []string {
	numericSlice, stringSlice := divideArrayByNumericAndStringPart(input)
	sort.Strings(stringSlice)
	sort.Slice(numericSlice, func(i, j int) bool {
		x, _ := strconv.Atoi(numericSlice[i][:len(numericSlice[i])-1])
		y, _ := strconv.Atoi(numericSlice[j][:len(numericSlice[j])-1])
		n1 := int64(x)
		n2 := int64(y)
		n1 *= m[numericSlice[i][len(numericSlice[i])-1]]
		n2 *= m[numericSlice[j][len(numericSlice[j])-1]]
		return n1 < n2
	})
	return mergeSlices(stringSlice, numericSlice)
}

func divideArrayByNumericAndStringPart(input []string) ([]string, []string) {
	numericSlice := make([]string, 0)
	stringSlice := make([]string, 0)
	for _, str := range input {
		fl := false
		for i := 0; i < len(str)-1; i++ {
			if !(str[i] >= '0' && str[i] <= '9') {
				fl = true
			}
		}

		if fl {
			stringSlice = append(stringSlice, str)
		} else {
			switch str[len(str)-1] {
			case 'B':
				numericSlice = append(numericSlice, str)
			case 'K':
				numericSlice = append(numericSlice, str)
			case 'M':
				numericSlice = append(numericSlice, str)
			case 'G':
				numericSlice = append(numericSlice, str)
			case 'T':
				numericSlice = append(numericSlice, str)
			case 'P':
				numericSlice = append(numericSlice, str)
			default:
				stringSlice = append(stringSlice, str)
			}
		}
	}
	return numericSlice, stringSlice
}
