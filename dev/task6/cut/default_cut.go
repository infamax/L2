package cut

import "strings"

func DefaultCut(words []string, field int, delimiter string) []string {
	res := make([]string, 0)
	for _, word := range words {
		separatedStrings := strings.Split(word, delimiter)
		res = append(res, separatedStrings[field-1])
	}
	return res
}
