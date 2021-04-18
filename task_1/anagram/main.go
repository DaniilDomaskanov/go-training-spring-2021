package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
  Description: An anagram is the result of rearranging the letters of a word to produce a new word (see wikipedia https://en.wikipedia.org/wiki/Anagram).
  Note: anagrams are case insensitive

  Complete the function to return true if the two arguments given are anagrams of each other; return false otherwise.

  Examples:
  "foefet" is an anagram of "toffee"
  "Buckethead" is an anagram of "DeathCubeK"
*/

func sortString(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func isAnagram(test, original string) bool {
	if strings.TrimSpace(original) == strings.TrimSpace(test) {
		return false
	}
	testOut := sortString(strings.ToLower(strings.Join(strings.Fields(test), "")))
	originalOut := sortString(strings.ToLower(strings.Join(strings.Fields(original), "")))
	return testOut == originalOut
}

func main() {
	fmt.Println(isAnagram("Buckethead", "DeathCubeK"))
	fmt.Println(isAnagram("foefet", "toffee"))
	fmt.Println(isAnagram("monkeys write", "New York Times"))
	fmt.Println(isAnagram("fdsfsdf", "New York Times"))
}
