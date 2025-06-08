package conditionals

import "fmt"

func ConditionalsTwo() {
	var bakiye float64 = 1000
	var cekilmekIstenen float64 = 1000

	if cekilmekIstenen > bakiye {
		fmt.Println("Bakiye yetersiz")
	} else if cekilmekIstenen == bakiye {
		fmt.Println("Bakiyeniz 0 olacak")
		bakiye = bakiye - cekilmekIstenen
	} else {
		fmt.Println("Paranızı buyrun")
		bakiye = bakiye - cekilmekIstenen
	}

	fmt.Printf("Bitti, hesaptaki para: %v", bakiye)
}
