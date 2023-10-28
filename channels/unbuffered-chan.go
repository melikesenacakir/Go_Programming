package channels

import "fmt"

//unbuffered channellarda her okuma yazma işlemi blockingtir,eşzamanlı 1 data yazıp okuyabiliyoruz
//bufferd olanlarda ise verdiğimiz sizea kadar bloklanmıyor sonrasında bloklanıyor
func Demo2() {
	unbufferedChan := make(chan int) //channel başlatma
	var unbufferedChan2 chan int     //channel tanımlama

	fmt.Println(unbufferedChan)
	fmt.Println(unbufferedChan2)

	go func(unbufcan <-chan int) {
		//veri gelene kadar bloklama yapar
		value := <-unbufcan
		fmt.Println(value)
	}(unbufferedChan)

	unbufferedChan <- 1
}
