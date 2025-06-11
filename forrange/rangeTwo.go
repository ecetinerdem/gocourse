package forrange

import "fmt"

func RangeTwo() {

	numbers := make([]int, 10)

	total := 0
	for i := range numbers {
		n := i + 1
		if n%2 != 0 {
			total += n
		}
	}
	fmt.Println(total)
}
