package error_handling

import "fmt"

func dogrula(i interface{}) {
	num, ok := i.(int)
	fmt.Println(num, ok) //hata varsa false yoksa tru verir ok
}

func Demo2() {
	var num interface{} = 5 //string veya başka tür girersek error
	dogrula(num)

	var num2 interface{} = "a"
	dogrula(num2)
}
