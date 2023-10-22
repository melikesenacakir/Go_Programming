package arrays

import "fmt"

func Demo2() {
	numbers := [5]int{20, 70, 90, 88, 65} //dizileri bu şekilde de tanımlayabiliriz
	enBuyuk := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if enBuyuk < numbers[i] {
			enBuyuk = numbers[i]
		}
	}
	fmt.Println("Dizideki en büyük sayı: " + fmt.Sprintf("%v", enBuyuk))
}
