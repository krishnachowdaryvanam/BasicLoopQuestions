// Write a Go function to calculate the sum of integers in a given array using a for loop.
package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	n := len(numbers)
	sum := 0
	for i := 0; i < n; i++ {
		sum = sum + numbers[i]
	}
	fmt.Println(sum)

}
