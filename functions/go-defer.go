package functions

import "fmt"

var mySlice = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
}

// DeferGotcha will have a number's scope shared
func DeferGotcha() {
	fmt.Println("DeferGotcha - Starting execution")
	for _, number := range mySlice {
		defer func() {
			fmt.Printf("   %s\n", number)
		}()
	}
}

// DeferGotchaImproved handle function will have you own number's scope
func DeferGotchaImproved() {
	fmt.Println("DeferGotchaImproved - Starting execution")
	handle := func(text string) {
		fmt.Printf("   %s\n", text)
	}

	for _, number := range mySlice {
		defer handle(number)
	}
}

// DeferGotchaImprovedAnonymousFunction defer will have your own number's scope
func DeferGotchaImprovedAnonymousFunction() {
	fmt.Println("DeferGotchaImprovedAnonymousFunction - Starting execution")
	for _, number := range mySlice {
		defer func(text string) {
			fmt.Printf("   %s\n", text)
		}(number)
	}
}
