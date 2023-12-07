package main

import "fmt"

func main() {
	n := 3
	a := checkPrimenumber(n)
	fmt.Println(a)
}

func checkPrimenumber(n int) bool {
	if n%2 == 0 {
		return true
	} else {
		return false
	}
}
