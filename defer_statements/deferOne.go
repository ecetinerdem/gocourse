package deferstatements

import "fmt"

func A() {
	fmt.Println("A works")

}

func B() {
	defer C()
	defer A()
	fmt.Println("B works")
}

func C() {
	fmt.Println("C works")
}
