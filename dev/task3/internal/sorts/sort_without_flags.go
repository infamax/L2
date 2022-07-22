package sorts

import (
	"golang.org/x/exp/constraints"
	"sort"
)

func SortWithoutKeys[T constraints.Ordered](input []T) []T {
	sort.Slice(input, func(i, j int) bool {
		return input[i] < input[j]
	})
	return input
}
