package channels

import "fmt"

func Demo6() {
	chan1 := make(chan int, 1)
	chan1 <- 8
	select { //select case ile veri dinleme
	case c1Val := <-chan1: //bir channella veri girişini bekletebiliyoruz
		fmt.Println(c1Val)
	}
}
