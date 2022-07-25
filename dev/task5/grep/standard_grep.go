package grep

import "regexp"

func StandardGrep(input []string, regularExpression string) ([]string, error) {
	reg, err := regexp.Compile(regularExpression)
	if err != nil {
		return []string{}, err
	}

	res := make([]string, 0)
	for _, str := range input {
		if reg.MatchString(str) {
			res = append(res, str)
		}
	}

	return res, nil
}
