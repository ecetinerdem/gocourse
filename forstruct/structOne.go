package forstruct

import "fmt"

func StructOne() {
	fmt.Println(product{"Computer", 500, "xyz", 0})
	fmt.Println(product{name: "Computer", unitPrice: 500})
}

type product struct {
	name         string
	unitPrice    float64
	brand        string
	discountRate int
}
