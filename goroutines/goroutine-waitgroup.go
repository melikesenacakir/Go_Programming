package goroutines

import (
	"fmt"
	"sync"
)

func Demo2() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		fmt.Println("test")
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("func dışı")
}
