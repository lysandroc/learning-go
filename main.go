package main

import (
	"fmt"

	f "./functions"
)

func main() {
	f.DeferGotcha()
	f.DeferGotchaImproved()
	f.DeferGotchaImprovedAnonymousFunction()

	nums := []int{1, 2, 3, 4}
	f.VariadicFunction(nums...)

	f.CreateMatrix()

	f.MapFriends()
	// f.MapDeleteItem()
	f.MapActivityOne()

	fmt.Print("\nhow iota works")
	const (
		_ = iota
		a
		b
		c = 1 << iota
		d
	)
	fmt.Println(a, b)
	fmt.Println(c, d)
	fmt.Println("end how iota works")
}
