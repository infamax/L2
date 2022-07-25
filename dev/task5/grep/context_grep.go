package grep

import (
	"regexp"
	"strings"
)

func ContextGrep(input []string, pattern string, after, before, context int,
	invert, ignoreCase bool) ([]string, error) {
	reg, err := regexp.Compile(pattern)
	if err != nil {
		return []string{}, nil
	}

	res := make([]string, 0)

	if !invert {
		if context != 0 {
			for i, word := range input {
				if !ignoreCase && reg.MatchString(word) {
					for j := i - 1; j >= max(i-context, 0); j-- {
						res = append(res, input[j])
					}
					res = append(res, word)
					for j := i + 1; j < min(i+context+1, len(input)); j++ {
						res = append(res, input[j])
					}
				} else if ignoreCase && reg.MatchString(strings.ToLower(word)) {
					for j := i - 1; j >= max(i-context, 0); j-- {
						res = append(res, input[j])
					}
					res = append(res, word)
					for j := i + 1; j < min(i+context+1, len(input)); j++ {
						res = append(res, input[j])
					}
				}
			}
		} else {
			for i, word := range input {
				if !ignoreCase && reg.MatchString(word) {
					for j := i - 1; j >= max(i-before, 0); j-- {
						res = append(res, input[j])
					}
					res = append(res, word)
					// [0, 1, 2] - words ["bar", "foo cyxs", "bar"]
					// before = 0
					for j := i + 1; j < min(i+after+1, len(input)); j++ {
						res = append(res, input[j])
					}
				} else if ignoreCase && reg.MatchString(strings.ToLower(word)) {
					for j := i - 1; j >= max(j-before, 0); j-- {
						res = append(res, input[j])
					}
					res = append(res, word)
					for j := i + 1; j < min(i+after+1, len(input)); j++ {
						res = append(res, input[j])
					}
				}
			}
		}
		return res, nil
	}

	if context != 0 {
		for i, word := range input {
			if !ignoreCase && !reg.MatchString(word) {
				for j := i - 1; j >= max(i-context, 0); j-- {
					res = append(res, input[j])
				}
				res = append(res, word)
				for j := i + 1; j < min(i+context+1, len(input)); j++ {
					res = append(res, input[j])
				}
			} else if ignoreCase && !reg.MatchString(strings.ToLower(word)) {
				for j := i - 1; j >= max(i-context, 0); j-- {
					res = append(res, input[j])
				}
				res = append(res, word)
				for j := i + 1; j < min(i+context+1, len(input)); j++ {
					res = append(res, input[j])
				}
			}
		}
	} else {
		for i, word := range input {
			if !ignoreCase && !reg.MatchString(word) {
				for j := i - 1; j >= max(i-before, 0); j-- {
					res = append(res, input[j])
				}
				res = append(res, word)
				for j := i + 1; j < min(i+after+1, len(input)); j++ {
					res = append(res, input[j])
				}
			} else if ignoreCase && !reg.MatchString(strings.ToLower(word)) {
				for j := i - 1; j >= max(i-before, 0); j-- {
					res = append(res, input[j])
				}
				res = append(res, word)
				for j := i + 1; j < min(i+after+1, len(input)); j++ {
					res = append(res, input[j])
				}
			}
		}
	}
	return res, nil
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
