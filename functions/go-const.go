package functions

import (
	"fmt"
)

const (
	first  = iota
	second = 2 << iota
	third  = iota
)

const (
	fourth = iota
	fifth
)

func Constants() {
	fmt.Println("first", first)
	fmt.Println("second", second)
	fmt.Println("third", third)
	fmt.Println("fourth", fourth)
	fmt.Println("fifth", fifth)
}
