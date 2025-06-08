package conditionals

import "fmt"

func Demo() {
	// üç adet int değişken
	//Ekrana en büyük olana yazdır

	var sayiBir int = 20
	var sayiIki int = 30
	var sayiUc int = 40

	if sayiBir > sayiIki {
		fmt.Printf("%v en büyük", sayiBir)
	} else if sayiIki > sayiUc {
		fmt.Printf("%v en büyük", sayiIki)
	} else {
		fmt.Printf("%v en büyük", sayiUc)
	}
}
