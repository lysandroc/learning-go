package functions

import "fmt"

func CreateMatrix() {
	fmt.Print("\nCreateMatrix - starting execution")
	sliceOfSlice:= [][][]int{
		[][]int{
			[]int{
				1,
				2,
			},
			[]int{
				3,
				4,
			},
		},
		[][]int{
			[]int{
				5,
				6,
			},
			[]int{
				7,
				8,
			},
		},
	}

	//fmt.Printf("   %v\n",sliceOfSlice)

	for _ , slice := range sliceOfSlice{
		fmt.Print("\n    [")
		for _, item := range slice {
			fmt.Printf("   		%v",item)
		}
		fmt.Print("    ]")
	}
	fmt.Print("\nCreateMatrix - ending execution ")
}