package loops

import "fmt"

func Workshop1() {
	num := 30
	i := 0
	fmt.Println("bir sayı söyleyin")
	fmt.Scanln(&i)
	counter := 1
	for i != num { //go da while yerine bu kullanılır
		if i > num {
			fmt.Println("Daha küçük bir sayı tahmin ediniz şuanda sayınız " + fmt.Sprintf("%v", i))
		} else if i < num {
			fmt.Println("Daha büyük bir sayı tahmin ediniz şuanda sayınız " + fmt.Sprintf("%v", i))
		}
		fmt.Scanln(&i)
		counter++
	}
	if counter < 4 {
		fmt.Println("Süper kısa sürede buldun!")
	} else if counter < 7 {
		fmt.Println("Güzel buldun!")
	} else {
		fmt.Println("Fena değil!")
	}
}
