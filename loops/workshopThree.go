package loops

import "fmt"

func DemoThree() {
	var sayi1 int = 1184
	var sayi2 int = 1210

	var karsiSayi1 int = 0
	var karsiSayi2 int = 0

	for i := 1; i <= sayi1/2; i++ {
		if sayi1%i == 0 {
			karsiSayi1 += i
		}
	}

	for i := 1; i <= sayi2/2; i++ {
		if sayi2%i == 0 {
			karsiSayi2 += i
		}
	}

	if sayi1 == karsiSayi2 && sayi2 == karsiSayi1 {
		fmt.Printf("%v ve %v arkadaş sayılardır", sayi1, sayi2)
	} else {
		fmt.Printf("%v ve %v arkadaş sayı değillerdir", sayi1, sayi2)
	}

}
