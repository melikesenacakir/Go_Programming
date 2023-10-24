package channels

func EvenNum(ciftSayiCn chan int) { //chan yazarak channel olduğunu belirtiriz
	toplam := 0
	for i := 0; i <= 10; i += 2 {
		toplam += i
	}
	ciftSayiCn <- toplam
}

func OddNum(tekSayiCn chan int) {
	toplam := 0
	for i := 1; i <= 10; i += 2 {
		toplam += i
	}
	tekSayiCn <- toplam
}
