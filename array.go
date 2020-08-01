package main

import "fmt"

func main() {
	number := 70
	guess := 50
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
