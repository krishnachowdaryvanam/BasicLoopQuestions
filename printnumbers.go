// Write a Go program to print numbers from 1 to 10 using a for loop.
package main

import "fmt"

// func main() {
// 	numbers := []int{1, 2, 3, 4, 5, 6}
// 	n := len(numbers)
// 	for i := 0; i < n; i++ {
// 		fmt.Println(i, numbers[i])
// 	}
// }

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	for i, a := range numbers {
		fmt.Println(i, a)
	}
}
