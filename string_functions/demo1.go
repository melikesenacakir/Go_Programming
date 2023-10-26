package string_functions

import (
	"fmt"
	s "strings" //strings yerine s harfi ile tanımlıycam demektir
)

func Demo1() {
	isim := "Melike"
	fmt.Println(s.Count(isim, "e"))
	fmt.Println(s.Contains(isim, "e"))
	sonuc := s.Index(isim, "e") // e harfini ilk gördüğü indexi döndürür
	if sonuc != -1 {
		fmt.Println("e harfi var")
	} else {
		fmt.Println("e harfi yok")
	}

}
