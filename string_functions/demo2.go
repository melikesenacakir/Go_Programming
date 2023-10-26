package string_functions

import (
	"fmt"
	s "strings"
)

func Demo2() {
	isim := "Melike"
	fmt.Println(s.HasPrefix(isim, "Mel")) //mel ile başluyor mu diye bakar
	fmt.Println(s.HasSuffix(isim, "en"))
	harfler := []string{"m", "e", "l", "i", "k", "e"}
	sonuc := s.Join(harfler, "*")
	fmt.Println(sonuc)
	fmt.Println(s.Replace(sonuc, "*", "+", -1)) //-1 ne görürsen gör değiştir demek onun dışında pozitif sayılar yazılan sayı kadar yer değiştir demek
}
