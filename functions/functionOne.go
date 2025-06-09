package functions

import "fmt"

func Sum(num1 int, num2 int) int {
	fmt.Printf("%v", num1+num2)
	return num1 + num2
}

func SayHi(name string) {
	fmt.Printf("Merhaba %v\n", name)
}
