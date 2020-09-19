package functions

import "fmt"

func VariadicFunction(nums ...int) {
	fmt.Println("VariadicFunction - Starting execution")
	fmt.Printf("   %d\n", nums)
	total := 0
	for _, num := range nums {
		total = total + num
	}
	fmt.Printf("   the sum of all numbers is %d\n", total)
	fmt.Println("VariadicFunction - Finishing execution")
}
