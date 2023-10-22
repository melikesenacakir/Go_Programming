package loops

import "fmt"

func Workshop3() {
	num1 := 0
	num2 := 0
	toplam1 := 0
	toplam2 := 0
	fmt.Println("Bir sayı giriniz")
	fmt.Scanln(&num1)
	fmt.Println("İkinci bir sayı giriniz")
	fmt.Scanln(&num2)
	if num1 < num2 {
		for i := 1; i < num2; i++ {
			if num1%i == 0 && i < num1 {
				toplam1 += i
			}
			if num2%i == 0 {
				toplam2 += i
			}
		}
	} else {
		for i := 1; i < num1; i++ {
			if num1%i == 0 {
				toplam1 += i
			}
			if num2%i == 0 && i <= num2 {
				toplam2 += i
			}
		}
	}

	if toplam1 == num2 && toplam2 == num1 {
		fmt.Println("Bu sayı arkadaş sayıdır")
	} else {
		fmt.Println("Bu sayı arkadaş sayı değildir")
	}
}
