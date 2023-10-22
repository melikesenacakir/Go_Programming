package conditions

import "fmt"

func Demo1() {
	account := 1000
	withdraw := 900
	if withdraw > account {
		fmt.Println("yeterli paranız bulunmamaktadır")
	} else {
		account -= withdraw
		fmt.Println("paranız başarıyla çekilmiştir kalan bakiye: " + fmt.Sprintf("%v", account)) //eğer ekrana değişkeni yazdırmak istersek Sprintf kullanırız
	}
}
