package deferstatements

import "fmt"

type product struct {
	productName string
	unitPrice   int
}

func (p product) save() {
	fmt.Println("Product kaydedildi")
}

func Demo1() {

	p := product{productName: "Laptop", unitPrice: 500}

	defer p.save()

	fmt.Println("islem basarılı")
}
