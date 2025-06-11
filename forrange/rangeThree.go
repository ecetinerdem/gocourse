package forrange

import "fmt"

func RangeThree() {

	// range for maps
	for k, v := range cities {
		fmt.Println(k)
		fmt.Printf(v)
	}
}
