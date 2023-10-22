package variables

import "fmt"

func Demo1() { //goda fonk adları pascal casedir.
	fmt.Println("Merhaba")
	var metin string = "hello world"
	fmt.Println(metin)
	var sayi int = 20
	fmt.Println(sayi)
	var kdv float64 = 0.1
	fmt.Println(100 + 100*kdv)
	kdv2 := 27 //veri tipini belirtmek istemezsek bu şekilde bir yazım yapılabilir
	fmt.Println(kdv2)
	fmt.Printf("veri tipi: %T \n", kdv2) //eğer verinin tipini yazdırmak istersek %T kullanırız
	//kdv2 := "hello" bu şekilde aynı değişkeni veri tipini değiştirerek kullanamayız
	kdv2 = 56
	fmt.Println(kdv2)

	var durum bool
	var text1 string = "hello"
	var text2 string = "hello"
	durum = text1 == text2 //go da bu şekilde karşılaştırma yapabiliyoruz
	fmt.Println(durum)
}
