package sorts

import (
	"golang.org/x/exp/constraints"
	"sort"
)

func SortUnique[T constraints.Ordered](input []T) []T {
	sort.Slice(input, func(i, j int) bool {
		return input[i] < input[j]
	})
	uniqueItems := make(map[T]struct{})
	res := make([]T, 0)
	for _, in := range input {
		_, ok := uniqueItems[in]
		if !ok {
			res = append(res, in)
		}
		uniqueItems[in] = struct{}{}
	}
	return res
}
