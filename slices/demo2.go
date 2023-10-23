package slices

import "fmt"

func Demo2() {
	cities := []string{"Ankara", "Çorum", "Kocaeli"}
	fmt.Println(cities)
	cities_copy := make([]string, len(cities))
	fmt.Println(cities_copy)
	copy(cities_copy, cities) //sehirlerin içindekini şehirler kopyaya kaydettik
	fmt.Println(cities_copy)
	cities_copy = append(cities_copy, "Adana")
	fmt.Println(cities_copy)
	fmt.Println(cities_copy[1:3]) //1den başlayıp 3e kadar yazdır demek(3 dahil değil)
	fmt.Println(cities_copy[1:])  //1den sonrakilerin hepsini yazıcaktır
	fmt.Println(cities_copy[:3])  //3e kadar olanların hepsini yazdıracaktır
}
