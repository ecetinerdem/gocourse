package arrays

import "fmt"

func ArraysThree() {
	numbers := [5]int{20, 30, 25, 10, 2}

	fmt.Println(numbers)

	var biggest int = numbers[0]

	for i := range len(numbers) {
		if numbers[i] > biggest {
			biggest = numbers[i]
		}
	}
	fmt.Println(biggest)
}
