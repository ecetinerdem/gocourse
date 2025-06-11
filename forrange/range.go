package forrange

import "fmt"

func RangeOne() {

	cities := []string{"Ankara", "Istanbul", "Izmir"}

	//index and item
	for _, city := range cities {
		fmt.Println(city)
	}

}
