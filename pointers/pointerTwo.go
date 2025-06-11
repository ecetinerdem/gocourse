package pointers

import "fmt"

func PointerTwo(numbers []int) {
	numbers[0] = 100
	fmt.Printf("Num in PointerTwo: %v\n", numbers[0])
}
