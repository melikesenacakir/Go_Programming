package error_handling

import (
	"errors"
	"fmt"
)

func enter_num(estimate int) (string, error) {
	num := 50
	if estimate < 1 || estimate > 100 {
		return "", errors.New("1-100 arasında bir sayı giriniz")
	}
	if estimate > num {
		return "Daha küçük sayı giriniz", nil
	}
	if estimate < num {
		return "Daha büyük sayı giriniz", nil
	}
	return "Doğru", nil
}

func Demo3() {
	mesaj, _ := enter_num(50)
	fmt.Println(mesaj)
	fmt.Println(enter_num(101))
	fmt.Println(enter_num(-10))
}
