package arrays

import "fmt"

func Demo1() {
	var numbers [5]int
	for i := 0; i < 5; i++ {
		fmt.Println("Sayi giriniz")
		fmt.Scanln(&numbers[i])
	}
	fmt.Println("\n\nDizi içerisindeki sayılarınız: ")
	for i := 0; i < 5; i++ {
		fmt.Println(numbers[i])
	}
}
