package main

import "fmt"

/*
 Description:
 Create a function that takes an integer as an argument and returns "Even" for even numbers or "Odd" for odd numbers.
*/

func evenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

func main() {
	fmt.Println(evenOrOdd(6))
	fmt.Println(evenOrOdd(7))
}
