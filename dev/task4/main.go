package main

import (
	"fmt"
	"sort"
	"strings"
)

func GroupAnagrams(words []string) map[string][]string {
	if len(words) == 0 {
		return nil
	}
	anagrams := getListAnagrams(makeSliceBytes(words))
	res := make(map[string][]string)

	for i := 0; i < len(anagrams); i++ {
		if len(anagrams[i]) == 1 {
			continue
		}
		res[anagrams[i][0]] = anagrams[i][1:]
	}
	return res
}

func makeSliceBytes(words []string) [][]rune {
	res := make([][]rune, len(words), len(words))
	for i := 0; i < len(words); i++ {
		res[i] = []rune(strings.ToLower(words[i]))
	}
	return res
}

func getListAnagrams(words [][]rune) [][]string {
	mapAnagrams := make(map[string][]string)
	for _, word := range words {
		tmp := make([]rune, 0, len(word))
		for i := 0; i < len(word); i++ {
			tmp = append(tmp, word[i])
		}
		sort.Slice(tmp, func(i, j int) bool {
			return tmp[i] < tmp[j]
		})
		mapAnagrams[string(tmp)] = append(mapAnagrams[string(tmp)], string(word))
	}

	res := make([][]string, 0, len(mapAnagrams))
	for _, val := range mapAnagrams {
		if len(val) > 1 {
			res = append(res, val)
		}
	}

	for _, val := range res {
		sort.Strings(val)
	}
	return res
}

func main() {
	words := []string{"пятак", "пятка", "тяпка",
		"листок", "слиток", "столик"}
	res := GroupAnagrams(words)
	fmt.Println(res)
}
