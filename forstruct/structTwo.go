package forstruct

import "fmt"

type customer struct {
	firstname string
	lastname  string
	age       int
}

func (customer) save() {
	fmt.Println("Added")
}

func StructTwo() {
	c := customer{"Engin", "Erdem", 35}
	c.save()
}
