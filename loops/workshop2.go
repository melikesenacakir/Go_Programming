package loops

import "fmt"

func Workshop2() {
	num := 0
	asal := 0
	fmt.Println("Bir sayı giriniz")
	fmt.Scanln(&num)
	if num == 1 {
		asal = 1
	}
	for i := 2; i < num; i++ {
		if num%i == 0 {
			asal = 1
			break
		}
	}
	if asal == 1 {
		fmt.Println("Bu sayı asal sayı değildir")
	} else {
		fmt.Println("Bu sayı asal sayıdır")
	}
}
