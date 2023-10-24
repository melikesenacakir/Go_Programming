package goroutines

import "fmt"

func EvenNum() {
	for i := 0; i < 10; i += 2 {
		fmt.Println("Çift sayı", i)
	}
}

func OddNum() {
	for i := 1; i < 10; i += 2 {
		fmt.Println("Tek sayı", i)
	}
}
