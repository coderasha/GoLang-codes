package main

import "fmt"

// Main function
func main() {
	for i := 0; i < 10; i++ {

		fmt.Println(i)
		if i == 4 {
			continue
		}
		// For loop breaks when the value of i = 3
		if i == 3 {
			break
		}

	}

}
