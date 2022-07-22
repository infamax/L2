package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	numericStringErr      = errors.New("string contains only numbers")
	numericFirstSymbolErr = errors.New("string first symbol digit")
)

const (
	noEscapeSeq = false
	escapeSeq   = true
)

func UnpackStringAdvanced(str string, escapeFl bool) (string, error) {
	if len(str) == 0 {
		return "", nil
	}

	if str[0] >= '1' && str[0] <= '9' {
		return "", numericFirstSymbolErr
	}

	if isNumericString(str) {
		return "", numericStringErr
	}

	if !escapeFl {
		i := 0
		var res strings.Builder
		for i < len(str) {
			j := i + 1
			for j < len(str) && (str[j] >= '1' && str[j] <= '9') {
				j++
			}
			num := str[i+1 : j]
			n, _ := strconv.ParseInt(num, 10, 64)
			var k int64 = 0

			res.WriteByte(str[i])

			for k = 1; k < n; k++ {
				res.WriteByte(str[i])
			}
			i = j
		}
		return res.String(), nil
	}

	var res strings.Builder
	fl := false
	i := 0
	for i < len(str) {
		if str[i] == '\\' && !fl {
			fl = true
			i++
			continue
		}

		if fl {
			j := i + 1
			for j < len(str) && (str[j] >= '1' && str[j] <= '9') {
				j++
			}
			num := str[i+1 : j]
			res.WriteByte(str[i])
			n, _ := strconv.ParseInt(num, 10, 64)
			var k int64 = 1
			for ; k < n; k++ {
				res.WriteByte(str[i])
			}
			i = j
			fl = false
			continue
		}
		res.WriteByte(str[i])
		i++
	}

	return res.String(), nil
}

func isNumericString(str string) bool {
	for i := 0; i < len(str); i++ {
		if !(str[i] >= '1' && str[i] <= '9') {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(UnpackStringAdvanced("qwe\\\\5", escapeSeq))
}
