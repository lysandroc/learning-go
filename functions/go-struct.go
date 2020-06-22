package functions

import (
	"fmt"
)

func StructExample() {
	type car struct {
		weight    float64
		sizeWheel int
		color     string
	}

	var carX car
	carX.color = "blue"
	carX.weight = 6.250
	carX.sizeWheel = 17
	fmt.Println(carX)

	CarY := car{
		color:     "red",
		sizeWheel: 20,
		weight:    7.920,
	}
	fmt.Println(CarY)

	//this create a new pointer
	CarZ := new(car)
	CarZ.color = "red"
	CarZ.sizeWheel = 16
	CarZ.weight = 5.500
	fmt.Println(CarZ)
}
