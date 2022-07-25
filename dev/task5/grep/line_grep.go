package grep

import (
	"fmt"
	"regexp"
	"strings"
)

func LineGrep(input []string, pattern string, invert bool, ignoreCase bool) ([]string, error) {
	reg, err := regexp.Compile(pattern)
	if err != nil {
		return []string{}, err
	}

	res := make([]string, 0)
	if !invert {
		for i, str := range input {
			if !ignoreCase && reg.MatchString(str) {
				res = append(res, fmt.Sprintf("%d: %s", i+1, str))
			} else if ignoreCase && reg.MatchString(strings.ToLower(str)) {
				res = append(res, fmt.Sprintf("%d: %s", i+1, str))
			}
		}
		return res, nil
	}

	for i, str := range input {
		if !ignoreCase && !reg.MatchString(str) {
			res = append(res, fmt.Sprintf("%d: %s", i+1, str))
		} else if ignoreCase && !reg.MatchString(strings.ToLower(str)) {
			res = append(res, fmt.Sprintf("%d: %s", i+1, str))
		}
	}

	return res, nil
}
