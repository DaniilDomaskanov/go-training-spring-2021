package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

/*
	given a string of words, return the length of the shortest word(s).
	String will never be empty and you do not need to account for different data types.
*/

func findShort(s string) int {
	total := strings.Split(s, " ")
	min := utf8.RuneCountInString(total[0])
	for _, value := range total {
		if min > utf8.RuneCountInString(value) {
			min = utf8.RuneCountInString(value)
		}
	}

	return min
}

func main() {
	fmt.Println(findShort("bitcoin take over the world maybe who knows perhaps"))
}
