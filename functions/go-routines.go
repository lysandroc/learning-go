package functions

import (
	"fmt"
	"sync"
	"time"
)

func GoRoutinesWaitGroupExample() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		defer waitGroup.Done()

		time.Sleep(5 * time.Second)
		fmt.Println("World")
	}()

	go func() {
		defer waitGroup.Done()
		fmt.Println("Hello")
	}()

	waitGroup.Wait()
}
