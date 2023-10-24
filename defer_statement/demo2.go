package defer_statement

import "fmt"

func Demo2(num int32) string {
	defer DeferFunc() //deferlar fonksiyonda en başa yazılmalı
	if num%2 == 0 {
		return "Even number"
	} else {
		return "Odd number"
	}
}
func Test() {
	result := Demo2(10)
	fmt.Println(result)
}
func DeferFunc() {
	fmt.Println("Finish")
}
