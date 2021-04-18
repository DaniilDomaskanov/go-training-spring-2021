package main

import (
	"fmt"
	"regexp"
)

/*
	Write a simple regex to validate a username. Allowed characters are:
	lowercase letters,
	numbers,
	underscore
	Length should be between 4 and 16 characters (both included).
	""
*/

func isUsername(username string) bool {
	matched, _ := regexp.Match("^[a-z0-9_]{4,16}$", []byte(username))
	return matched
}

func main() {
	fmt.Println(isUsername("danik123"))
}
