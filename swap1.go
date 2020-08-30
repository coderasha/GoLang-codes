package main

import "fmt"

func swap(a, b string) (string, string) {

	return b, a
}
func main() {
	x := "ghug"
	y := "hh"
	swap(x, y)
	fmt.Println("", y)
}
