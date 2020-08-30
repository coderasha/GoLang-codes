

import (
	"fmt"
)

func main() {
	odd := [6]int{2, 4, 6, 8, 10, 12}
	var s []int = odd[1:]
	fmt.Println(s)
}
