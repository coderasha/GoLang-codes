package main

import "fmt"

func main() {

	a := "Mahesh"
	b := "Kumar"

	var temp string
	temp = a
	a = b
	b = temp

	fmt.Println("", a, b)

}
