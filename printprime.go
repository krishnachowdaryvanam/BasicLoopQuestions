// Create a program that prints all even numbers from 1 to 20 using a for loop and an if statement.
package main

func main() {
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			println(i)
		}
	}
}
