package for_range

import "fmt"

func Demo1() {
	cities := []string{"Ankara", "Çorum", "Kocaeli"}
	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}
	for i, city := range cities { //foreach döngüsüdür ve i yi eğer indexi kullanıcaksak yazarız yoksa yerine _ yazarız
		fmt.Print(i + 1)
		fmt.Println(city)
	}
}
