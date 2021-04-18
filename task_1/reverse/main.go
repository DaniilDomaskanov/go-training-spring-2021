package main

import "fmt"

/**
 * Description: Complete the solution so that it reverses the string passed into it.
 * Example: "world"  =>  "dlrow"
 */

func reverse(word string) string {
	reverseWord := []rune(word)

	for i, j := 0, len(reverseWord)-1; i < j; i, j = i+1, j-1 {
		reverseWord[i], reverseWord[j] = reverseWord[j], reverseWord[i]
	}

	return string(reverseWord)
}

func main() {
	fmt.Println(reverse("world"))
	fmt.Println(reverse("Hello"))
	fmt.Println(reverse("Geeks"))
}
