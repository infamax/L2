package sorts

import (
	"fmt"
	"sort"
	"strconv"
)

func SortByNumericValue(input []string) []string {
	// 1 step: divide slice in two : []string and []float64
	strs, floats := divideSlice(input)
	// 2 step: sort slices
	sort.Strings(strs)
	sort.Float64s(floats)
	// 3 step: merge slice in one
	input = make([]string, 0, len(strs)+len(floats))
	for _, val := range strs {
		input = append(input, val)
	}
	for _, val := range floats {
		input = append(input, fmt.Sprintf("%f", val))
	}
	return input
}

func divideSlice(input []string) ([]string, []float64) {
	strs := make([]string, 0)
	floats := make([]float64, 0)
	for _, in := range input {
		n, err := strconv.ParseFloat(in, 64)
		if err != nil {
			strs = append(strs, in)
		} else {
			floats = append(floats, n)
		}
	}
	return strs, floats
}
