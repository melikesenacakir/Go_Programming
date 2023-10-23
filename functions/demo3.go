package functions

func SumVariadic(numbers ...int) int { //bu şekilde parametre sayısı belirsiz fonksiyon oluştururuz
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return sum
}
