package main

import (
	"fmt"
)

func main() {
	items := []string{"Item1", "Item2", "Item3"}
	items = make([]string, 5, 10)

	fmt.Printf("items have \nlength:%d \ncapacity:%d \n", len(items), cap(items))

	for index, item := range items {
		fmt.Printf("index: %d has item of %s \n", index, item)
	}
}
