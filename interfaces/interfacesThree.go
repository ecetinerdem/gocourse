package interfaces

import "fmt"

func verify(i interface{}) {
	sayi, ok := i.(int)
	fmt.Println(sayi, ok)
}

func Demo3() {
	var num interface{} = 5
	var noNum interface{} = "Engin"
	verify(5)
	verify(num)
	verify(noNum)
}
