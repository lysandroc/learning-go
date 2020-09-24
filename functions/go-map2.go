package functions

import "fmt"

func MapFriends() {
	fmt.Print("\n\n Map example - starting execution \n\n")

	friends := map[string]int{
		"lysandro": 123456,
		"carioca":  654321,
	}

	fmt.Printf("friends %v \n", friends)

	if value, ok := friends["cariocaaa"]; ok {
		fmt.Println("lysandro: ", value)
	} else {
		fmt.Println("this item could be found")
	}

	fmt.Print("\n\n Map example - ending execution")
}
