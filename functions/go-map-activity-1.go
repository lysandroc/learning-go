package functions

import "fmt"

func MapActivityOne() {
	fmt.Print("\n\n Map activity one - starting execution \n\n")

	activities := [5]int{1, 2, 3, 4, 5}
	for _, value := range activities {
		fmt.Printf("\n value: %v", value)
	}
	fmt.Printf("\n Type of map %T", activities)
	fmt.Print("\n\n Map activity one - ending execution")
}
