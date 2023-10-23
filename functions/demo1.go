package functions

import "fmt"

func Sum(a int, b int) int { // bu şekilde sağına yazarak fonksiyonun void olmadığını belirtebiliriz
	return a + b
}

func SayHi() {
	fmt.Println("Hello")
}
