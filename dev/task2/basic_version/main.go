package main

import (
	"fmt"
	"strconv"
	"strings"
)

func UnpackString(str string) (string, bool) {
	if len(str) == 0 {
		return "", true
	}

	if str[0] >= '1' && str[0] <= '9' {
		return "", false
	}

	if isNumericString(str) {
		return "", false
	}

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
	return res.String(), true
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
	fmt.Println(UnpackString("2a"))
}
