package channels

import "fmt"

//eğer 1den fazla olursa select ile dinledğimiz channellar, bunlardan birisi random olarak seçilip basılır ikisi birden basılmaz,eğer 2 tane kullanmak istersek şu şekilde yazarız

func Demo7() {
	chan1 := make(chan int, 1)
	chan1 <- 11

	chan2 := make(chan int, 1)
	chan2 <- 8

	var done bool
	for !done {
		select {
		case c1Val := <-chan1:
			fmt.Println(c1Val)
		case c2Val := <-chan2:
			fmt.Println(c2Val)
		default:
			done = true
		}
	}
}
