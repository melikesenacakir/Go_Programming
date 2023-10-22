package conditions

import "fmt"

func Workshop1() {
	var1 := 20
	var2 := 89
	var3 := 66
	if var1 > var2 && var1 > var3 {
		fmt.Println(fmt.Sprintf("%v", var1) + "  aralarındaki en büyük sayıdır")
	} else if var2 > var1 && var2 > var3 {
		fmt.Println(fmt.Sprintf("%v", var2) + "  aralarındaki en büyük sayıdır")
	} else {
		fmt.Println(fmt.Sprintf("%v", var3) + "  aralarındaki en büyük sayıdır")
	}
}
