package main

import "fmt"

func main() {

	var a [10]int
	var n int
	fmt.Printf("Enter the length%d", n)
	fmt.scanf(&n)
	a[0] = 0
	a[1] = 1

	for i := 2; i < n; i++ {
		a[i] = a[i-2] + a[i-1]

	}

	fmt.Println(a[0:10])
}
