package functions

func Dortislem(num1 int, num2 int) (int, int, int, float64) { //multiple return fonksiyonları böyle tanımlarız
	sum := num1 + num2
	sub := num1 - num2
	mul := num1 * num2
	div := float64(num1) / float64(num2)
	return sum, sub, mul, div

}
