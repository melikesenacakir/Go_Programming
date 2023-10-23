package arrays

import "fmt"

func Workshop1() {
	var numbers [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println("diziye sayı giriniz")
			fmt.Scanln(&numbers[i][j])

		}
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("numbers[%v][%v]=", i, j)
			fmt.Println(numbers[i][j])

		}
	}

}
