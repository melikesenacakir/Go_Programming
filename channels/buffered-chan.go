package channels

import "fmt"

func Demo3() {
	bufferedChan := make(chan int, 5)

	bufferedChan <- 1
	bufferedChan <- 2
	bufferedChan <- 3 //bunu da max 5e kadar yazarız çünkü size 5 tanımlandı

	fmt.Println(<-bufferedChan)
	fmt.Println(<-bufferedChan)
	fmt.Println(<-bufferedChan)
	//fmt.Println(<-bufferedChan) tanımlanandan fazla yazdırırsak error verir

}
