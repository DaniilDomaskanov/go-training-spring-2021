package main

import (
	"fmt"
	"math"
)

/*
	Define a function that takes one integer argument and returns logical value true or false depending on if the integer is a prime.

	Per Wikipedia, a prime number (or a prime) is a natural number greater than 1 that has no positive divisors other than 1 and itself.

	Requirements
	You can assume you will be given an integer input.
	You can not assume that the integer will be only positive. You may be given negative numbers as well (or 0).
	NOTE on performance: There are no fancy optimizations required, but still the most trivial solutions might time out. Numbers go up to 2^31 (or similar, depends on language version). Looping all the way up to n, or n/2, will be too slow.

	Example
	isPrime(1)  -> false
	isPrime(2)  -> true
	isPrime(-1) -> false
*/

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}

	maxDivider := int(math.Sqrt(float64(n)) + 1)

	for i := 2; i < maxDivider; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPrime(1))
	fmt.Println(isPrime(2))
	fmt.Println(isPrime(-1))
	fmt.Println(isPrime(3))
}
