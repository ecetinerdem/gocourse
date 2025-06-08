package arrays

import "fmt"

func ArraysTwo() {
	var cities [5]string
	cities[0] = "Ankara"
	cities[1] = "Istanbul"
	cities[2] = "Izmir"
	cities[3] = "Adana"
	cities[4] = "Diyarbakır"

	fmt.Println(cities)

	for i := range len(cities) {
		fmt.Println(cities[i])
	}

}
