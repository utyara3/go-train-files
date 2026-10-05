package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	retMap := map[string]int{}

	words := strings.SplitSeq(s, " ")
	for word := range words {
		if _, ok := retMap[word]; ok {
			retMap[word] += 1
		} else {
			retMap[word] = 1
		}
	}

	return retMap
}

func main() {
	s := "привет привет пока пока пока тест"

	fmt.Printf("%s\nWords count: %v\n", s, WordCount(s))
}
