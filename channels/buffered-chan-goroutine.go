package channels

import (
	"fmt"
	"time"
)

func Demo4() {
	bufferedChan := make(chan int, 3)
	go func(bufChan chan int) {
		for {
			val := <-bufChan
			fmt.Println(val)
		}
	}(bufferedChan)

	bufferedChan <- 1
	bufferedChan <- 2
	bufferedChan <- 3
	time.Sleep(time.Second)
	//output non-determİnistic
}
