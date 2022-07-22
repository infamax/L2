package sorts

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

func SortByColumn(input []string, column int) ([]string, error) {
	if column < 0 {
		return nil, errors.New(fmt.Sprintf("sort: invalid number at field start: "+
			"invalid count at start of '%d'", column))
	}

	if column == 0 {
		return nil, errors.New("sort: field number is zero: invalid field specification ‘0’")
	}

	words := make([][]string, 0, len(input))

	for _, str := range input {
		strs := strings.Split(str, " ")
		words = append(words, strs)
	}

	res := make([]string, 0)

	if column > len(words[0]) {
		sort.Slice(words, func(i, j int) bool {
			return words[i][0] < words[j][0]
		})
		for i := 0; i < len(words); i++ {
			res = append(res, strings.Join(words[0], " "))
		}
		return res, nil
	}

	sort.Slice(words, func(i, j int) bool {
		return words[i][column-1] < words[j][column-1]
	})

	for i := 0; i < len(words); i++ {
		res = append(res, strings.Join(words[i], " "))
	}

	return res, nil
}
