// Reverse String

package main

import "fmt"

func main() {
	integer := 123789
	str := "chikki"
	i := reverseInt(integer)
	s := reverseString(str)
	fmt.Println(i)
	fmt.Println(s)
}

func reverseInt(n int) int {
	res := 0
	for n > 0 {
		remainder := n % 10
		res = res*10 + remainder
		n = n / 10
	}
	return res
}

func reverseString(n string) string {
	s := ""
	for i := len(n) - 1; i >= 0; i-- {
		s += string(n[i])
	}
	return s
}
