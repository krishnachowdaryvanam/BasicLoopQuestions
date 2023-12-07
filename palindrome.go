package main

import "fmt"

func main() {
	n := -121
	original := n
	reverse := 0
	for n > 0 {
		remainder := n % 10
		reverse = reverse*10 + remainder
		n /= 10
	}
	if original == reverse {
		fmt.Println("The Given is palindrome")
	} else {
		fmt.Println("The Given is not palindrome")
	}
}
