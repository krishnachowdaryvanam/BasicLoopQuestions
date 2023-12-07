// Iterate over a slice of strings in Go and print each element.
package main

import "fmt"

func main() {
	str := []string{"chikki", "kittu", "krishna", "supreetha"}
	n := len(str)
	for i := 0; i < n; i++ {
		fmt.Println(str[i])
	}
}
