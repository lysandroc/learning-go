package functions

import "fmt"

func MapRange() {
	fmt.Print("\n\n Map Range example - starting execution \n\n")

	maps := map[string]string{
		"lysandro": "five avenue",
		"carioca":  "fix avenue",
	}

	for key, value := range maps {
		fmt.Printf("\n name: %s, street: %s\n", key, value)
	}

	delete(maps, "carioca")

	all_streets := ""
	for _, value := range maps {
		all_streets += value
	}
	fmt.Printf("\n streets: %s", all_streets)
	fmt.Print("\n\n Map Range example - ending execution")
}
