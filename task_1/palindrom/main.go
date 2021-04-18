package main

import (
	"fmt"
	"strings"
)

/*
 Description: A palindrome is a word, phrase, number, or other
 sequence of characters which reads the same backward or forward.
 This includes capital letters, punctuation, and word dividers.

 Implement a function that checks if something is a palindrome.

 Examples:
 isPalindrome("anna")   ==> true
 isPalindrome("walter") ==> false
 isPalindrome("12321")    ==> true
 isPalindrome("123456")   ==> false
*/

func isPalindrome(str string) bool {
	str = strings.ToLower(str)
	palindrome := []rune(str)

	for i, j := 0, len(palindrome)-1; i < j; i, j = i+1, j-1 {
		palindrome[i], palindrome[j] = palindrome[j], palindrome[i]
	}

	return str == string(palindrome)
}

func main() {
	fmt.Println(isPalindrome("Anna"))
	fmt.Println(isPalindrome("walter"))
	fmt.Println(isPalindrome("12321"))
	fmt.Println(isPalindrome("123456"))
}
