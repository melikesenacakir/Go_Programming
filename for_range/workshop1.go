package for_range

import "fmt"

func Workshop1() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	toplam := 0
	for _, num := range numbers {
		if num%2 != 0 {
			toplam += num
		}
	}
	fmt.Println("Tek sayıların toplamı", toplam)
}
