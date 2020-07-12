package main

import (
	f "./functions"
)

func main() {
	f.DeferGotcha()
	f.DeferGotchaImproved()
	f.DeferGotchaImprovedAnonymousFunction()
	//Learning variadic function
	nums := []int{1, 2, 3}
	f.VariadicFunction(nums...)
}
