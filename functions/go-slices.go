package functions

import (
	"fmt"
)

func SliceExamples() {
	items := []string{"Item1", "Item2", "Item3"}
	items = make([]string, 5, 10)

	fmt.Printf("items have \nlength:%d \ncapacity:%d \n", len(items), cap(items))

	for index, item := range items {
		fmt.Printf("index: %d has item of %s \n", index, item)
	}

	mySlice := []int{1, 2, 3, 4, 5}
	myNewSlice := []int{10, 20, 30}
	mySlice = append(mySlice, myNewSlice...)
	fmt.Println(mySlice)
}
