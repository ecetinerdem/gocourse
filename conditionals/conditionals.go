package conditionals

import "fmt"

func Conditionals() {
	var bakiye float64 = 1000
	var cekilmekIstenen float64 = 500

	if cekilmekIstenen > bakiye {
		fmt.Println("Bakiye yetersiz")
	} else {
		fmt.Println("Paranızı buyrun")
		bakiye = bakiye - cekilmekIstenen
	}
	fmt.Printf("Bitti, hesaptaki para: %v", bakiye)
}
