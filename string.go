// Write a Go function that takes a string as input and prints each character
// on a new line using a for loop.
package main

import "fmt"

func main() {
	str := "kittu"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n", str[i])
	}
}
