package main

import (
	"golesson/restful"
)

//"golesson/conditions"
//"golesson/variables"

func main() {
	//variables.Demo1()
	//conditions.Demo1()
	//conditions.Workshop1()
	//loops.Demo1()
	//loops.Workshop3()
	//arrays.Workshop1()
	//slices.Demo2()
	//toplam := functions.Sum(3, 4)
	//fmt.Println(toplam)
	//functions.SayHi()

	//sum, sub, mul, div := functions.Dortislem(9, 4)
	//fmt.Println("Toplam", sum)
	//fmt.Println("Çıkarım", sub)
	//fmt.Println("Çarpım", mul)

	//fmt.Println("Bölüm", div)
	//sonuc := functions.SumVariadic(3, 4, 5)
	//fmt.Println(sonuc)
	//bu da diğer yolu sayilar:=[]int{2,3,4,5,9} fmt.Println(functions.SumVariadic(sayilar...)) 3 nokta koyarak variadic olduğunu belirtiriz
	//maps.Demo1()
	//for_range.Demo1()
	//for_range.Workshop1()
	//sayi := 4
	//pointers.Demo1(&sayi)
	//fmt.Println(sayi)
	//structs.Demo2()
	//go goroutines.EvenNum()
	//go goroutines.OddNum()
	//time.Sleep(5 * time.Second) gerçekte bunun yerine channel kullanırız

	//ciftSayiCn := make(chan int)
	//tekSayiCn := make(chan int)
	//go channels.EvenNum(ciftSayiCn)
	//go channels.OddNum(tekSayiCn)
	//ciftSayiToplam, tekSayiToplam := <-ciftSayiCn, <-tekSayiCn //bu channel okumadır
	//carpim := ciftSayiToplam * tekSayiToplam
	//fmt.Println("Çarpım:", carpim) asenkron çalışma
	//interfaces.Demo2()
	//defer_statement.B()
	//defer_statement.Demo3()
	//error_handling.Demo3()
	//fmt.Println(error_handling.Demo4(333))
	//string_functions.Demo2()
	restful.Demo2()
}
