package loops

import "fmt"

func Demo() {

	var sayi int = 29

	var tahminEdilenSayi = 0
	var tahmin int = 1

	for {
		fmt.Println("1 ve 100 arasında bir sayi girin")
		fmt.Scanln(&tahminEdilenSayi)

		if tahminEdilenSayi > 100 || tahminEdilenSayi <= 0 {
			fmt.Println("1 ve 100 arasında bir sayi girin")
		} else if tahminEdilenSayi > sayi {
			fmt.Println("Tahmin ettiğin sayı büyük")
		} else if tahminEdilenSayi < sayi {
			fmt.Println("Tahmin ettiğin sayı  küçük")
		} else {
			fmt.Println("Tebrikler")
			break
		}
		tahmin++
	}

	if tahmin <= 3 {
		fmt.Printf("%v, tahmin de bildiniz. Süper", tahmin)
	} else if tahmin <= 6 {
		fmt.Printf("%v, tahmin de bildiniz, Güzel", tahmin)
	} else {
		fmt.Printf("%v, tahmin de bildiniz, Fena değil", tahmin)
	}
}
