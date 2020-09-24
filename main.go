package main

import (
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
}
