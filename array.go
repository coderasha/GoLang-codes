package main

import "fmt"

func main() {
	number := 70
	guess := 70
	if guess < number {
		fmt.Println("low")
	}
	if guess > number {
		fmt.Println("high")
	}
	if guess == number {
		fmt.Println("same")
	}
}
