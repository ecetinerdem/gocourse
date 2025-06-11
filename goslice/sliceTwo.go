package goslice

import "fmt"

func SlicesTwo() {

	cities := []string{"Ankara", "Istanbul", "Izmir"}

	copyCities := make([]string, len(cities))

	copy(copyCities, cities)

	copyCities = append(copyCities, "Muğla")

	fmt.Println(copyCities)

	fmt.Println(cities[0:2])
}
