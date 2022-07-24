package cut

import "strings"

func DeleteNotSeparatedStrings(words []string, delimiter string) []string {
	res := make([]string, 0)
	for _, word := range words {
		if strings.Contains(word, delimiter) {
			res = append(res, word)
		}
	}
	return res
}
