package defer_statement

import "fmt"

func A() { //defer demek kodun diğer kısımlarında hata olsada defer olan fonksiyon mutlaka çalışacak demektir
	fmt.Println("A fonksiyonu çalıştı")
}
func C() {
	fmt.Println("C fonksiyonu çalıştı")
}

func B() {
	defer A()
	defer C() //Lifo yapıdadırlar(stack)
	fmt.Println("B fonksiyonu çalıştı")
}
