package sorts

import (
	"sort"
	"strings"
)

func SortIgnoreLeadingBlanks(input []string) []string {
	sort.Slice(input, func(i, j int) bool {
		s1 := strings.TrimSpace(input[i])
		s2 := strings.TrimSpace(input[j])
		return s1 < s2
	})
	return input
}
