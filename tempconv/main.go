package main

import (
	"fmt"

	"github.com/lysandroc/tempconv"
)

func main() {
	celsiusFromKelvin := tempconv.KToC(10)
	fmt.Sprint(celsiusFromKelvin)
}
