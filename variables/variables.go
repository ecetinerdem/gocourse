package variables

import "fmt"

func Variables() {
	fmt.Println("Hello, World!")

	var kdv int = 64
	fmt.Println(kdv)

	var kdv2 float32 = 0.2
	fmt.Println(kdv2)

	kdv3 := 25.2
	fmt.Println(kdv3)

	fmt.Printf("Type of: %T\n", kdv3)

	var condition bool = false
	fmt.Println(condition)
}
