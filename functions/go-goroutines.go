package functions

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func GoRoutinesExample() {
	f("directly")

	go func(param string) {
		fmt.Println(param)
	}("goroutine between")

	go f("through goroutines")

	time.Sleep(time.Second)
	fmt.Println("done")
}
