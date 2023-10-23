package slices

import "fmt"

func Demo1() {
	names := make([]string, 3)
	for i := 0; i < len(names); i++ {
		fmt.Println("isim giriniz")
		fmt.Scanln(&names[i])
	}
	names = append(names, "melike") //diziye yeni eleman eklemesi yapabiliyorz ama bunu sadece diziyi make fonksiyonu ile oluşturduysak yapabiliriz
	fmt.Println(names)
}
