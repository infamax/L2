package sorts

func CheckSortedSlice(input []string) bool {
	if len(input) <= 1 {
		return true
	}

	for i := 1; i < len(input); i++ {
		if input[i] < input[i-1] {
			return false
		}
	}
	return true
}
