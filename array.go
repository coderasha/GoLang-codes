package main

import "fmt"

func main() {
	var n [10]int /* n is an array of 10 integers */
	var i int

	/* initialize elements of array n to 0 */
	for i := 0; i < 10; i++ {
		n[i] = i + 100 /* set element at location i to i + 100 */

		/* output each array element's value */

		fmt.Printf("Element[%d] = %d\n", i, n[i])
	}
}
