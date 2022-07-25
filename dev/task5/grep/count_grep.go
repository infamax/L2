package grep

import "regexp"

func CountGrep(input []string, pattern string, invert bool) (int, error) {
	reg, err := regexp.Compile(pattern)
	if err != nil {
		return 0, err
	}

	count := 0

	if !invert {
		for _, str := range input {
			if reg.MatchString(str) {
				count++
			}
		}
		return count, nil
	}

	for _, str := range input {
		if !reg.MatchString(str) {
			count++
		}
	}
	return count, nil
}
