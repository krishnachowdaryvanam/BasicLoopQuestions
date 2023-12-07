// Write a Go program that prints the numbers from 1 to 100. But for multiples of 3,
// print "Fizz" instead of the number, and for multiples of 5, print "Buzz".
// For numbers that are multiples of both 3 and 5, print "FizzBuzz".
package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println(i, "Fizz")
		} else if i%5 == 0 {
			fmt.Println(i, "Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
