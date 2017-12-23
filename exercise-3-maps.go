package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {

	wordMap := make(map[string]int)
	
	for _, word := range strings.Fields(s) {
		
		if count, ok := wordMap[word]; ok {
			wordMap[word] = count + 1
		} else {
			wordMap[word] = 1
		}
	}
	
	return wordMap
}

func main() {
	wc.Test(WordCount)
}
