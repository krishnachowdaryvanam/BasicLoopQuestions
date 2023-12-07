package main

import "fmt"

func fact(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result = result * i
	}
	return result
}

func main() {
	n := 4
	a := fact(n)
	fmt.Println(a)
}
