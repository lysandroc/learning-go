package functions

import "fmt"

func fib(n int) chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for i, j := 0, 1; i < n; i, j = i+j, i {
			out <- i
		}
	}()

	return out
}

func CreateChannel() {
	for x := range fib(1000000) {
		fmt.Println(x)
	}
}
