package main

import (
	f "./functions"
)

func main() {
	f.DeferGotcha()
	f.DeferGotchaImproved()
	f.DeferGotchaImprovedAnonymousFunction()

	f.CreateMatrix()

	f.MapFriends()
	// f.MapDeleteItem()
	f.MapActivityOne()
}
