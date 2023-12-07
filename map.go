// Create a Go program that iterates over a map and prints both the keys and values.
package main

import "fmt"

func main() {
	n := map[int]string{
		1: "name",
		2: "DOB",
		3: "Age",
	}
	for i, a := range n {
		fmt.Println(i, a)
	}
}
