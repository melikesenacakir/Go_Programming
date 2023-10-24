package structs

import "fmt"

func Demo1() {
	fmt.Println(product{"Bilgisayar", 9005.66, "XYZ"})
	//eğer tüm alanları doldurmadan bilgi girmek istersek
	fmt.Println(product{name: "Bilgisayar2", price: 9005.66})
}

type product struct {
	name  string
	price float64
	brand string
}
