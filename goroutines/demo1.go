package goroutines

import "fmt"

func EvenNums() {
	for i := 0; i <= 10; i += 2 {
		fmt.Println(i)
	}
}

func OddNums() {
	for i := 1; i <= 10; i += 2 {
		fmt.Println(i)
	}

}
