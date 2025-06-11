package pointers

import "fmt"

func PointerOne(number *int) {
	*number = *number + 1
	fmt.Printf("Num in PointerOne: %v\n", *number)
}
