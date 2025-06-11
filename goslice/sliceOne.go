package goslice

import "fmt"

func SlicesOne() {
	names := make([]string, 3)

	names[0] = "Engin"
	names[1] = "Engin"
	names[2] = "Engin"
	names = append(names, "Cetin")

	fmt.Println(names)
}
