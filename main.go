package main

import "fmt"

func main() {
	fmt.Println(countPairs([]int{1, 3, 2, 2, 3, 4}, 5)) // Output: 2
	fmt.Println(countPairs([]int{1, 1, 1, 1}, 2))       // Output: 1
	fmt.Println(countPairs([]int{1, 2, 3, 4, 5}, 7))    // Output: 2

}

func countPairs(numbers []int, targetNumber int) int {
	var res int

	alreadySeen := make(map[int]bool) // create a variable to make sure a number is seen/viewed only once
	combination := make(map[int]bool) // create a variable to make sure 2 number combo is only used once

	for _, number := range numbers {
		// find out the number to add
		secondNumber := targetNumber - number

		// add the number to variable "res" after the number is already viewed
		if alreadySeen[secondNumber] && !combination[secondNumber*number] {
			res++
			combination[secondNumber*number] = true // set the combination to true so it will only be used once
		}

		alreadySeen[number] = true
	}

	return res
}
