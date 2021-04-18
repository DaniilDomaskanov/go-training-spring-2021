package main

import (
	"fmt"
	"sort"
)

/*
 You are given an array of n+1 integers 1 through n. In addition there is a single duplicate integer.

 The array is unsorted.

 An example valid array would be [3, 2, 5, 1, 3, 4]. It has the integers 1 through 5 and 3 is duplicated. [1, 2, 4, 5, 5] would not be valid as it is missing 3.

 You should return the duplicate value as a single integer.
*/

func isValidArray(numbers []int) bool {
	sort.Ints(numbers)
	for j := 1; j < len(numbers); j++ {
		index := sort.SearchInts(numbers, j)
		if index == len(numbers) || j != numbers[index] {
			return false
		}
	}
	return true
}

// return -1 if array doesn't have any duplicates
func getDuplicate(numbers []int) int {
	if !isValidArray(numbers) {
		return -1
	}
	duplicateNumber := -1
	for i := 1; i < len(numbers); i++ {
		if numbers[i-1] == numbers[i] {
			duplicateNumber = numbers[i]
		}
	}
	return duplicateNumber
}

func main() {
	fmt.Println(getDuplicate([]int{3, 2, 1, 5, 4, 4}))
	fmt.Println(getDuplicate([]int{1, 3, 4, 5, 2, 3}))
}
