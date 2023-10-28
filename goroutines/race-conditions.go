package goroutines

import (
	"fmt"
	"sync"
)

func RaceExampleFixed() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	val := 0
	lock := sync.Mutex{} //mutex memory adresine eşzamanlı olarak sadece 1 go rutininin erişebileceğini garanti altına alır
	go func() {
		for i := 0; i < 10; i++ {
			lock.Lock() //bu memory adresine sadece bu lock yazılan erişebilir
			val++
			lock.Unlock()
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(val)

}
