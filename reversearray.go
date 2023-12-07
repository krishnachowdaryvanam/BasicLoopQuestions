// Reserve a Array

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	strings := []string{"apple", "banana", "cherry", "date"}
	a := reverseArr(arr)
	b := reverseString(strings)
	fmt.Println("Reverse int array:", a)
	fmt.Println("Reverse string array:", b)

}

func reverseString(n []string) []string {
	s := make([]string, 0, len(n))
	for i := len(n) - 1; i >= 0; i-- {
		s = append(s, n[i])
	}
	return s
}

func reverseArr(n []int) []int {
	s := make([]int, 0, len(n))
	for i := len(n) - 1; i >= 0; i-- {
		s = append(s, n[i])
	}
	return s
}
