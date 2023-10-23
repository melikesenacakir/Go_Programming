package maps //sözlük yapısı
import "fmt"

func Demo1() {
	dict := make(map[string]string) //[] içi key type dışı ise value type
	dict["table"] = "Masa"
	dict["book"] = "kitap"
	dict["pencil"] = "kalem"
	fmt.Println(dict)
	delete(dict, "book")
	fmt.Println(dict)
	deger, varMi := dict["table"] //key value değerlerini atamak istersek
	fmt.Println("Lisede olma durumu", varMi)
	fmt.Println("Degeri", deger)
}
