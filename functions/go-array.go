package functions

import "fmt"

var x [5]int
var y [6]int

func CreateArray() {

	x = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("x %v of %T \n", x, x)

	z := [5]int{0, 1, 2, 3, 4}
	fmt.Printf("z %v of %T \n", z, z)
}
