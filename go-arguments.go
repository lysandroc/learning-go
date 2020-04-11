package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	//first way
	fmt.Println(os.Args[1:])

	//second way
	var args, space string
	for _, argument := range os.Args[1:] {
		args += space + argument
		space = " "
	}
	fmt.Println(args)

	//third way
	var str, separator string
	for i := 1; i < len(os.Args); i++ {
		str += separator + os.Args[i]
		separator = " "
	}
	fmt.Println(str)

	//fourth way
	fmt.Println(strings.Join(os.Args[1:], " "))
}
