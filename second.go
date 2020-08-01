package main

import "fmt"

const PI = 3.14

var SI int = 10

func main() {
	const GFG = "GeeksforGeeks"
	fmt.Println("Hello", GFG)

	fmt.Println("Happy", PI, "Day")

	const Correct = true
	fmt.Println("Go rules?", Correct)
	fmt.Println("The number is :&d", SI)
}
