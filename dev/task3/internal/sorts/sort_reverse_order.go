package sorts

import (
	"golang.org/x/exp/constraints"
	"sort"
)

func SortReverseOrder[T constraints.Ordered](input []T) ([]T, bool) {
	if len(input) <= 1 {
		return input, true
	}

	fl := true
	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			fl = false
		}
	}

	if fl {
		return input, fl
	}

	sort.Slice(input, func(i, j int) bool {
		return input[i] > input[j]
	})

	return input, fl
}
